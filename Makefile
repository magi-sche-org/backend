.PHONY: help env build build-local up down logs ps test generate migrate
.DEFAULT_GOAL := help

env: ## Initialize project
	cp example.env .env

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t geekcamp-vol11-team30/backend:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

init-local: ## Initialize local development
	@make env
	@make build-local
	@make up
	@make migrate

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

exec: ## Execute command in container
	docker compose exec -it app bash

exec-db: ## Execute command in container
	docker compose exec -it db mysql -umysql -pmysql magische

test: ## Execute tests
	docker compose exec -it app go test -race -shuffle=on ./...  -coverprofile=coverage.out

generate: ## Go generate
	rm -rf ./db/models
	docker compose exec -it app go generate ./...

migrate: ## Execute migration
	docker compose exec -it app goose -dir ./db/migrations mysql "mysql:mysql@tcp(db:3306)/magische?parseTime=true&multiStatements=true" up
migrate-down: ## Execute migration
	docker compose exec -it app goose -dir ./db/migrations mysql "mysql:mysql@tcp(db:3306)/magische?parseTime=true&multiStatements=true" down

migrate-remote: ## Execute migration on remote
	docker build --target migrate -t geekcamp-vol11-team30/backend-migrate:${DOCKER_TAG} .
	docker run geekcamp-vol11-team30/backend-migrate:${DOCKER_TAG}

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
