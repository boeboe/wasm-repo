# Copyright (c) Tetrate, Inc 2021 All Rights Reserved.
.DEFAULT_GOAL := help
.PHONY: help
help: Makefile ## This help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n"} \
			/^[.a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36mmake %-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# Variables
SERVER_DIR  := api
BINARY_NAME := wasm-repo-server
BUILD_DIR   := build

DB_TYPE=postgres
# DB_TYPE=mysql
# DB_TYPE=sqlite

POSTGRES_CONTAINER := wasm-postgres
POSTGRES_USER      := gorm
POSTGRES_PASSWORD  := gorm
POSTGRES_DB        := gorm
POSTGRES_PORT      := 9920

MYSQL_CONTAINER := wasm-mysql
MYSQL_USER      := gorm
MYSQL_PASSWORD  := gorm
MYSQL_DB        := gorm
MYSQL_PORT      := 3306

.PHONY: build
build: ## Build the server
	@echo "Building the server..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SERVER_DIR)/main.go

.PHONY: run
run: build ## Run the server
	@echo "Running the server..."
	@/bin/bash -c "DB_TYPE=${DB_TYPE} ./$(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)/*

.PHONY: start-postgres
start-postgres: ## Start postgres db
	@echo "Starting PostgreSQL container..."
	@docker run --name ${POSTGRES_CONTAINER} \
							--env POSTGRES_USER=${POSTGRES_USER} \
							--env POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
							--env POSTGRES_DB=${POSTGRES_DB} \
							--publish ${POSTGRES_PORT}:5432 \
							--detach postgres:latest
	@echo "PostgreSQL started on port 9920."

.PHONY: stop-postgres
stop-postgres: ## Stop postgres db
	@echo "Stopping PostgreSQL container..."
	@docker stop ${POSTGRES_CONTAINER}
	@docker rm ${POSTGRES_CONTAINER}
	@echo "PostgreSQL container stopped and removed."

.PHONY: start-mysql
start-mysql: ## Start MySQL
	@echo "Starting MySQL container..."
	@docker run --name ${MYSQL_CONTAINER} \
							--env MYSQL_ROOT_PASSWORD=${MYSQL_PASSWORD} \
							--env MYSQL_DATABASE=${MYSQL_DB} \
							--env MYSQL_USER=${MYSQL_USER} \
							--env MYSQL_PASSWORD=${MYSQL_PASSWORD} \
							-p ${MYSQL_PORT}:3306 \
							--detach mysql:latest
	@echo "MySQL started on port 3306."

.PHONY: stop-mysql
stop-mysql: ## Stop MySQL
	@echo "Stopping MySQL container..."
	@docker stop ${MYSQL_CONTAINER}
	@docker rm ${MYSQL_CONTAINER}
	@echo "MySQL container stopped and removed."

