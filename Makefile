.PHONY: docker-build docker-run-in-memory docker-run-postgres unit-test

unit-test:
	mockgen -source=internal/url/repository.go  -destination=internal/url/mock/mock_repository.go -package=mock
	mockgen -source=internal/url/usecase.go  -destination=internal/url/mock/mock_usecase.go -package=mock
	go test ./...

docker-build:
	docker build -t ozon/base -f docker/base/Dockerfile .
	docker build -t ozon/url -f docker/microservices/url/Dockerfile .
	docker build -t ozon/api -f docker/microservices/api/Dockerfile .
	docker image rm ozon/base

docker-run-in-memory:
	docker compose -f deployment/in_memory/docker-compose.yml up

docker-run-postgres:
	docker compose -f deployment/postgres/docker-compose.yml up

