# Copyright (c) Tetrate, Inc 2021 All Rights Reserved.
.DEFAULT_GOAL := help
.PHONY: help
help: Makefile ## This help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n"} \
			/^[.a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36mmake %-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)


# Variables
SERVER_DIR  :=api
BINARY_NAME :=wasm-repo-server
BUILD_DIR   :=build

POSTGRES_CONTAINER := wasm-postgres
POSTGRES_USER      := gorm
POSTGRES_PASSWORD  := gorm
POSTGRES_DB        := gorm
POSTGRES_PORT      := 9920

.PHONY: build
build: ## Build the server
	@echo "Building the server..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SERVER_DIR)/main.go

.PHONY: run
run: build ## Run the server
	@echo "Running the server..."
	@./$(BUILD_DIR)/$(BINARY_NAME)


.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

.PHONY: start-db
start-db: ## Start postgres db
	@echo "Starting PostgreSQL container..."
	@docker run --name ${POSTGRES_CONTAINER} \
							--env POSTGRES_USER=${POSTGRES_USER} \
							--env POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
							--env POSTGRES_DB=${POSTGRES_DB} \
							--publish ${POSTGRES_PORT}:5432 \
							--detach postgres:latest
	@echo "PostgreSQL started on port 9920."

.PHONY: stop-db
stop-db: ## Stop postgres db
	@echo "Stopping PostgreSQL container..."
	@docker stop ${POSTGRES_CONTAINER}
	@docker rm ${POSTGRES_CONTAINER}
	@echo "PostgreSQL container stopped and removed."
