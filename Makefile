#!make
.SILENT:
.DEFAULT_GOAL := help

compose-up: ## Run server and client by docker-compose
	docker-compose up

compose-down: ## Stop server and client by docker-compose
	docker-compose down

compose-clean: ## Stop and remove containers, networks, images, and volumes by docker-compose
	docker-compose down --rmi all --volumes --remove-orphans

test: ## Run tests
	go test ./...

start-server: ## Run only server
	TCP_ADDRESS=0.0.0.0:8081 TCP_KEEP_ALIVE=20s TCP_DEADLINE=20s POW_DIFFICULTY=4 LOG_LEVEL=debug go run cmd/server/main.go

start-client: ## Run only client
	SERVER_ADDR=0.0.0.0:8081 go run cmd/client/main.go


deps: ## Download dependencies
	go mod download && go mod tidy
