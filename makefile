ROOT:=$(shell name)
LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54322 dbname=edication-bot user=edication-bot-user password=edication-bot-password"

BIN_SENDER := ./bin
PATH_STATIC_DATA := ./static

.PHONY: run/server
run/bot:
		go run ./cmd/main.go ${PATH_STATIC_DATA}
.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: migrate-up
migrate-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: migrate-down
migrate-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

.PHONY: create-migration
create-migration:
	@cd migrations; \
	goose create $(ARGS) sql
.PHONY: run/db
run/db:
	docker-compose up pg-edication-bot-db

