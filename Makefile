test.integration:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes
test.integration.host:
	go test -tags integration ./internal/tests/integration/... -v
test.unit:
	go test ./internal/tests/unit/... -v