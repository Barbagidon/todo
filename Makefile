include .env
export

PROJECT_ROOT := $(shell pwd)
export PROJECT_ROOT

.PHONY: \
	postgres-up \
	postgres-stop \
	postgres-restart \
	postgres-logs \
	postgres-ps \
	compose-down \
	postgres-clean

postgres-up:
	docker compose up -d todoapp-postgres

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