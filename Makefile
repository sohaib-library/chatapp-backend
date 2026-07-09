include .env

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATION_DIR=./database/migration

.PHONY: migrate-up migrate-down migrate-status create-migration run docker-up docker-down docker-logs

migrate-up:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" status

create-migration:
	goose -dir $(MIGRATION_DIR) create $(name) sql

run:
	go run main.go