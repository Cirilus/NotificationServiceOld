package keycloak

import "gopkg.in/square/go-jose.v2/jwt"

type Token struct {
	Jti               string                 `json:"jti,omitempty"`
	Exp               int64                  `json:"exp"`
	Nbf               int64                  `json:"nbf"`
	Iat               int64                  `json:"iat"`
	Iss               string                 `json:"iss"`
	Sub               string                 `json:"sub"` // id
	Typ               string                 `json:"typ"`
	Azp               string                 `json:"azp,omitempty"`
	Nonce             string                 `json:"nonce,omitempty"`
	AuthTime          int64                  `json:"auth_time,omitempty"`
	SessionState      string                 `json:"session_state,omitempty"`
	Acr               string                 `json:"acr,omitempty"`
	ClientSession     string                 `json:"client_session,omitempty"`
	AllowedOrigins    []string               `json:"allowed-origins,omitempty"`
	ResourceAccess    map[string]ServiceRole `json:"resource_access,omitempty"`
	Name              string                 `json:"name"`
	PreferredUsername string                 `json:"preferred_username"`
	GivenName         string                 `json:"given_name,omitempty"`
	FamilyName        string                 `json:"family_name,omitempty"`
	Email             string                 `json:"email,omitempty"`
	RealmAccess       ServiceRole            `json:"realm_access,omitempty"`
	CustomClaims      interface{}            `json:"custom_claims,omitempty"`
}

type ServiceRole struct {
	Roles []string `json:"roles"`
}

type Certs struct {
	Keys []KeyEntry `json:"keys"`
}

type KeyEntry struct {
	Kid string   `json:"kid"`
	Kty string   `json:"kty"`
	Alg string   `json:"alg"`
	Use string   `json:"use"`
	Crv string   `json:"crv"`
	X   string   `json:"x"`
	Y   string   `json:"y"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5C []string `json:"x5c"`
}

type ClaimMapperFunc func(jsonWebToken *jwt.JSONWebToken, keyCloakToken *Token) error

type Config struct {
	Url                string
	Realm              string
	FullCertsPath      *string
	CustomClaimsMapper ClaimMapperFunc
}
