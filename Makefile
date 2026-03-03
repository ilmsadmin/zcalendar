# ──────────────────────────────────────────────────────────────────────
# Zplus Calendar – Makefile
# ──────────────────────────────────────────────────────────────────────

# Variables
APP_NAME     := zplus-calendar
BACKEND_DIR  := backend
FRONTEND_DIR := frontend
BINARY       := $(BACKEND_DIR)/bin/$(APP_NAME)
GO           := go
GOFLAGS      := -v

# Docker
COMPOSE      := docker compose
COMPOSE_FILE := docker-compose.yml

# ──────────────────────────────────────────────────────────────────────
# Development
# ──────────────────────────────────────────────────────────────────────

.PHONY: dev
dev: ## Run backend with hot-reload (Air)
	cd $(BACKEND_DIR) && air

.PHONY: dev-frontend
dev-frontend: ## Run frontend dev server
	cd $(FRONTEND_DIR) && npm run dev

.PHONY: dev-all
dev-all: ## Run both backend and frontend
	@make -j2 dev dev-frontend

# ──────────────────────────────────────────────────────────────────────
# Build
# ──────────────────────────────────────────────────────────────────────

.PHONY: build
build: ## Build the Go binary
	cd $(BACKEND_DIR) && $(GO) build $(GOFLAGS) -o bin/$(APP_NAME) ./cmd/server

.PHONY: build-frontend
build-frontend: ## Build the Next.js frontend
	cd $(FRONTEND_DIR) && npm run build

.PHONY: build-all
build-all: build build-frontend ## Build everything

.PHONY: run
run: build ## Build and run the backend
	cd $(BACKEND_DIR) && ./bin/$(APP_NAME)

# ──────────────────────────────────────────────────────────────────────
# Testing
# ──────────────────────────────────────────────────────────────────────

.PHONY: test
test: ## Run all Go tests
	cd $(BACKEND_DIR) && $(GO) test ./... -v

.PHONY: test-lunar
test-lunar: ## Run only lunar engine tests
	cd $(BACKEND_DIR) && $(GO) test ./internal/lunar/... -v

.PHONY: test-cover
test-cover: ## Run tests with coverage
	cd $(BACKEND_DIR) && $(GO) test ./... -coverprofile=coverage.out
	cd $(BACKEND_DIR) && $(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: $(BACKEND_DIR)/coverage.html"

.PHONY: test-race
test-race: ## Run tests with race detector
	cd $(BACKEND_DIR) && $(GO) test ./... -race -v

.PHONY: bench
bench: ## Run benchmarks
	cd $(BACKEND_DIR) && $(GO) test ./internal/lunar/... -bench=. -benchmem

# ──────────────────────────────────────────────────────────────────────
# Code Quality
# ──────────────────────────────────────────────────────────────────────

.PHONY: lint
lint: ## Run golangci-lint
	cd $(BACKEND_DIR) && golangci-lint run ./...

.PHONY: fmt
fmt: ## Format Go source files
	cd $(BACKEND_DIR) && $(GO) fmt ./...

.PHONY: vet
vet: ## Run go vet
	cd $(BACKEND_DIR) && $(GO) vet ./...

.PHONY: tidy
tidy: ## Tidy Go modules
	cd $(BACKEND_DIR) && $(GO) mod tidy

.PHONY: check
check: fmt vet lint test ## Run all checks (format, vet, lint, test)

# ──────────────────────────────────────────────────────────────────────
# Docker
# ──────────────────────────────────────────────────────────────────────

.PHONY: docker-up
docker-up: ## Start all services with Docker Compose
	$(COMPOSE) -f $(COMPOSE_FILE) up -d

.PHONY: docker-down
docker-down: ## Stop all Docker services
	$(COMPOSE) -f $(COMPOSE_FILE) down

.PHONY: docker-rebuild
docker-rebuild: ## Rebuild and restart services
	$(COMPOSE) -f $(COMPOSE_FILE) up -d --build

.PHONY: docker-logs
docker-logs: ## Tail logs from all services
	$(COMPOSE) -f $(COMPOSE_FILE) logs -f

.PHONY: docker-ps
docker-ps: ## Show running containers
	$(COMPOSE) -f $(COMPOSE_FILE) ps

.PHONY: docker-db
docker-db: ## Start only infrastructure (postgres, mongo, redis)
	$(COMPOSE) -f $(COMPOSE_FILE) up -d postgres mongo redis

.PHONY: docker-clean
docker-clean: ## Stop services and remove volumes
	$(COMPOSE) -f $(COMPOSE_FILE) down -v --remove-orphans

# ──────────────────────────────────────────────────────────────────────
# Database
# ──────────────────────────────────────────────────────────────────────

.PHONY: db-migrate
db-migrate: ## Run database migrations (via GORM AutoMigrate)
	cd $(BACKEND_DIR) && $(GO) run ./cmd/server --migrate

.PHONY: db-seed
db-seed: ## Seed the database with initial data
	cd $(BACKEND_DIR) && $(GO) run ./cmd/seed

.PHONY: db-reset
db-reset: ## Drop and recreate the database
	@echo "⚠️  This will destroy all data!"
	@read -p "Continue? [y/N] " confirm && [ "$$confirm" = "y" ] || exit 1
	$(COMPOSE) -f $(COMPOSE_FILE) exec postgres psql -U zplus -c "DROP DATABASE IF EXISTS zplus_calendar;"
	$(COMPOSE) -f $(COMPOSE_FILE) exec postgres psql -U zplus -c "CREATE DATABASE zplus_calendar;"
	@echo "Database reset complete."

.PHONY: db-shell
db-shell: ## Open psql shell
	$(COMPOSE) -f $(COMPOSE_FILE) exec postgres psql -U zplus -d zplus_calendar

.PHONY: mongo-shell
mongo-shell: ## Open mongosh
	$(COMPOSE) -f $(COMPOSE_FILE) exec mongo mongosh -u zplus -p zplus_secret_2024 zplus_calendar

.PHONY: redis-cli
redis-cli: ## Open redis-cli
	$(COMPOSE) -f $(COMPOSE_FILE) exec redis redis-cli

# ──────────────────────────────────────────────────────────────────────
# Swagger / Docs
# ──────────────────────────────────────────────────────────────────────

.PHONY: swagger
swagger: ## Generate Swagger documentation
	cd $(BACKEND_DIR) && swag init -g cmd/server/main.go -o docs/swagger

# ──────────────────────────────────────────────────────────────────────
# Clean
# ──────────────────────────────────────────────────────────────────────

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(BACKEND_DIR)/bin
	rm -f  $(BACKEND_DIR)/coverage.out $(BACKEND_DIR)/coverage.html
	rm -rf $(FRONTEND_DIR)/.next $(FRONTEND_DIR)/node_modules

# ──────────────────────────────────────────────────────────────────────
# Help
# ──────────────────────────────────────────────────────────────────────

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
