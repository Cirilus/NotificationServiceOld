package keycloak

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"
	"math/big"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// VarianceTimer controls the max runtime of Auth() and AuthChain() middleware
var VarianceTimer = 30000 * time.Millisecond
var publicKeyCache = cache.New(8*time.Hour, 8*time.Hour)

// TokenContainer stores all relevant token information
type TokenContainer struct {
	Token         *oauth2.Token
	KeyCloakToken *Token
}

func extractToken(r *http.Request) (*oauth2.Token, error) {
	hdr := r.Header.Get("Authorization")
	if hdr == "" {
		return nil, errors.New("No authorization header")
	}

	th := strings.Split(hdr, " ")
	if len(th) != 2 {
		return nil, errors.New("Incomplete authorization header")
	}

	return &oauth2.Token{AccessToken: th[1], TokenType: th[0]}, nil
}

func GetTokenContainer(token *oauth2.Token, config Config) (*TokenContainer, error) {

	keyCloakToken, err := decodeToken(token, config)
	if err != nil {
		return nil, err
	}

	return &TokenContainer{
		Token: &oauth2.Token{
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
		},
		KeyCloakToken: keyCloakToken,
	}, nil
}

func getPublicKey(keyId string, config Config) (interface{}, error) {

	keyEntry, err := getPublicKeyFromCacheOrBackend(keyId, config)
	if err != nil {
		return nil, err
	}
	if strings.ToUpper(keyEntry.Kty) == "RSA" {
		n, _ := base64.RawURLEncoding.DecodeString(keyEntry.N)
		bigN := new(big.Int)
		bigN.SetBytes(n)
		e, _ := base64.RawURLEncoding.DecodeString(keyEntry.E)
		bigE := new(big.Int)
		bigE.SetBytes(e)
		return &rsa.PublicKey{bigN, int(bigE.Int64())}, nil
	} else if strings.ToUpper(keyEntry.Kty) == "EC" {
		x, _ := base64.RawURLEncoding.DecodeString(keyEntry.X)
		bigX := new(big.Int)
		bigX.SetBytes(x)
		y, _ := base64.RawURLEncoding.DecodeString(keyEntry.Y)
		bigY := new(big.Int)
		bigY.SetBytes(y)

		var curve elliptic.Curve
		crv := strings.ToUpper(keyEntry.Crv)
		switch crv {
		case "P-224":
			curve = elliptic.P224()
		case "P-256":
			curve = elliptic.P256()
		case "P-384":
			curve = elliptic.P384()
		case "P-521":
			curve = elliptic.P521()
		default:
			return nil, errors.New("EC curve algorithm not supported " + keyEntry.Kty)
		}

		return &ecdsa.PublicKey{
			Curve: curve,
			X:     bigX,
			Y:     bigY,
		}, nil
	}

	return nil, errors.New("no support for keys of type " + keyEntry.Kty)
}

func getPublicKeyFromCacheOrBackend(keyId string, config Config) (KeyEntry, error) {
	entry, exists := publicKeyCache.Get(keyId)
	if exists {
		return entry.(KeyEntry), nil
	}

	u, err := url.Parse(config.Url)
	if err != nil {
		return KeyEntry{}, err
	}

	if config.FullCertsPath != nil {
		u.Path = *config.FullCertsPath
	} else {
		u.Path = path.Join(u.Path, "realms", config.Realm, "protocol/openid-connect/certs")
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return KeyEntry{}, err
	}
	defer resp.Body.Close()
	var certs Certs
	err = json.NewDecoder(resp.Body).Decode(&certs)

	if err != nil {
		return KeyEntry{}, err
	}

	if err != nil {
		return KeyEntry{}, err
	}

	for _, keyIdFromServer := range certs.Keys {
		if keyIdFromServer.Kid == keyId {
			publicKeyCache.Set(keyId, keyIdFromServer, cache.DefaultExpiration)
			return keyIdFromServer, nil
		}
	}

	return KeyEntry{}, errors.New("No public key found with kid " + keyId + " found")
}

func decodeToken(token *oauth2.Token, config Config) (*Token, error) {
	keyCloakToken := Token{}

	var err error
	parsedJWT, err := jwt.ParseSigned(token.AccessToken)
	if err != nil {
		logrus.Errorf("[Gin-OAuth] jwt not decodable: %s", err)
		return nil, err
	}
	key, err := getPublicKey(parsedJWT.Headers[0].KeyID, config)
	if err != nil {
		logrus.Errorf("Failed to get publickey %v", err)
		return nil, err
	}

	err = parsedJWT.Claims(key, &keyCloakToken)
	if err != nil {
		logrus.Errorf("Failed to get claims JWT:%+v", err)
		return nil, err
	}

	if config.CustomClaimsMapper != nil {
		err = config.CustomClaimsMapper(parsedJWT, &keyCloakToken)
		if err != nil {
			logrus.Errorf("Failed to get custom claims JWT:%+v", err)
			return nil, err
		}
	}

	return &keyCloakToken, nil
}

func isExpired(token *Token) bool {
	if token.Exp == 0 {
		return false
	}
	now := time.Now()
	fromUnixTimestamp := time.Unix(token.Exp, 0)
	return now.After(fromUnixTimestamp)
}

func getTokenContainer(ctx *gin.Context, config Config) (*TokenContainer, bool) {
	var oauthToken *oauth2.Token
	var tc *TokenContainer
	var err error

	if oauthToken, err = extractToken(ctx.Request); err != nil {
		logrus.Errorf("[Gin-OAuth] Can not extract oauth2.Token, caused by: %s", err)
		return nil, false
	}
	if !oauthToken.Valid() {
		logrus.Infof("[Gin-OAuth] Invalid Token - nil or expired")
		return nil, false
	}

	if tc, err = GetTokenContainer(oauthToken, config); err != nil {
		logrus.Errorf("[Gin-OAuth] Can not extract TokenContainer, caused by: %s", err)
		return nil, false
	}

	if isExpired(tc.KeyCloakToken) {
		logrus.Errorf("[Gin-OAuth] Keycloak Token has expired")
		return nil, false
	}

	return tc, true
}

func (t *TokenContainer) Valid() bool {
	if t.Token == nil {
		return false
	}
	return t.Token.Valid()
}

func Auth(accessCheckFunction AccessCheckFunction, endpoints Config) gin.HandlerFunc {
	return authChain(endpoints, accessCheckFunction)
}

func authChain(config Config, accessCheckFunctions ...AccessCheckFunction) gin.HandlerFunc {
	// middleware
	return func(ctx *gin.Context) {
		t := time.Now()
		varianceControl := make(chan bool, 1)

		go func() {
			tokenContainer, ok := getTokenContainer(ctx, config)
			if !ok {
				_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("no token in context"))
				varianceControl <- false
				return
			}

			if !tokenContainer.Valid() {
				_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid Token"))
				varianceControl <- false
				return
			}
			ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "token", tokenContainer.KeyCloakToken))
			ctx.Next()
			for _, fn := range accessCheckFunctions {
				if fn(tokenContainer, ctx) {
					varianceControl <- true
					return
				}
			}
			_ = ctx.AbortWithError(http.StatusForbidden, errors.New("Access to the Resource is forbidden"))
			varianceControl <- false
			return
		}()

		select {
		case ok := <-varianceControl:
			if !ok {
				logrus.Infof("[Gin-OAuth] %12v %s access not allowed", time.Since(t), ctx.Request.URL.Path)
				return
			}
		case <-time.After(VarianceTimer):
			_ = ctx.AbortWithError(http.StatusGatewayTimeout, errors.New("Authorization check overtime"))
			logrus.Infof("[Gin-OAuth] %12v %s overtime", time.Since(t), ctx.Request.URL.Path)
			return
		}

		logrus.Infof("[Gin-OAuth] %12v %s access allowed", time.Since(t), ctx.Request.URL.Path)
	}
}

func RequestLogger(keys []string, contentKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.Request
		c.Next()
		err := c.Errors
		if request.Method != "GET" && err == nil {
			data, e := c.Get(contentKey)
			if e != false { //key is non existent
				values := make([]string, 0)
				for _, key := range keys {
					val, keyPresent := c.Get(key)
					if keyPresent {
						values = append(values, val.(string))
					}
				}
				logrus.Infof("[Gin-OAuth] Request: %+v for %s", data, strings.Join(values, "-"))
			}
		}
	}
}
