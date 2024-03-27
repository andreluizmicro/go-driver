COMMAND_DOCKER=docker compose

docker-build:
	docker compose build

docker-up:
	docker compose up -d

docker-down: 
	docker compose down

docker-bash:
	$(COMMAND_DOCKER) exec app bash

go-test:
	go test ./...  

go-test-coverage:
	@mkdir -p coverage
	go test -coverprofile=./coverage/coverage.out ./...
	go tool cover -html=./coverage/coverage.out
