include .env
export

PROJECT_ROOT := $(shell pwd)
export PROJECT_ROOT

MIGRATE_DB_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable
MIGRATE_CMD := docker compose run --rm todoapp-postgres-migrate -path=/migrations -database "$(MIGRATE_DB_URL)"

.PHONY: \
	postgres-up \
	postgres-forward-up \
	postgres-forward-stop \
	postgres-stop \
	postgres-restart \
	postgres-logs \
	postgres-ps \
	compose-down \
	postgres-clean \
	migrate-create \
	migrate-up \
	migrate-down

postgres-up:
	docker compose up -d todoapp-postgres

postgres-forward-up:
	docker compose up -d todoapp-postgres port-forwarder

postgres-forward-stop:
	docker compose stop port-forwarder

postgres-stop:
	docker compose stop todoapp-postgres

postgres-restart:
	docker compose restart todoapp-postgres

postgres-logs:
	docker compose logs -f todoapp-postgres

postgres-ps:
	docker compose ps

compose-down:
	docker compose down

postgres-clean:
	@read -p "Delete postgres data? [y/N] " confirm; \
	if [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ]; then \
		docker compose down; \
		rm -rf $(PROJECT_ROOT)/out/pgdata; \
		echo "Postgres data deleted"; \
	else \
		echo "Cancelled"; \
	fi

migrate-create:
	@test -n "$(name)" || (echo "Error: variable 'name' is not set. Usage: make migrate-create name=create_users_table" && exit 1)
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(name)"

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down $(or $(steps),1)
