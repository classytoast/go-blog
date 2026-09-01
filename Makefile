include .env
export


export PROJECT_ROOT=$(shell pwd)


env-up:
	@docker compose up -d blog-postgres

env-down:
	@docker compose down blog-postgres port-forwarder

env-cleanup:
	@read -p "Очистить все volume файлы окружения? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down blog-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi


env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder


# В будущих проектах лучше использовать простую последовательную нумерацию
# миграций через -seq, а не timestamp-версии. Например: -seq "$(seq)"
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm blog-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		"$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action. Пример action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm blog-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@blog-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

migrate-up:
	@$(MAKE) migrate-action action=up

migrate-down:
	@$(MAKE) migrate-action action=down


blogapp-run:
	@set -e; \
	export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	if [ "$(tidy)" = "true" ]; then \
		go mod tidy || exit 1; \
	fi; \
	go run ${PROJECT_ROOT}/cmd/blogapp/main.go

