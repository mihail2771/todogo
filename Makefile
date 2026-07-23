include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Очистить все volume - ОПАСНО !!! [y/N]: " ans;\
	if [ "$$ans" = "y" ]; then\
		docker compose down todoapp-postgres port-forwarding &&\
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы очищены"; \
	else \
		echo "Очистка отменена"; \
	fi

env-port-forvard:
	@docker compose up -d port-forwarding

env-port-close:
	@docker compose down port-forwarding


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует seq"; \
		exit 1; \
	fi; \
	mkdir -p migrations; \
	docker compose run --rm todoapp-postpores-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)" 

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует action"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

logs-cleanup:
	@read -p "Очистить все логи - ОПАСНО !!! [y/N]: " ans;\
	if [ "$$ans" = "y" ]; then\
		rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Файлы очищены"; \
	else \
		echo "Очистка отменена"; \
	fi

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go

