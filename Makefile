# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/server/main.go

# Run the application
run:
	@go run server.go

#Run first with a migration name: Create migration state 
migrate-diff:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then \
		echo "Error: You must provide a name for the migration."; \
		exit 1; \
	fi
	@echo "Running atlas migrate diff $(word 2,$(MAKECMDGOALS)) --env dev"
	@./scripts/atlas.sh diff $(word 2,$(MAKECMDGOALS)) --env dev
%:
	@:

#apply migration
migrate-apply:
	@echo "Running atlas migrate apply"
	@./scripts/atlas-apply.sh apply

#apply migration
migrate-down:
	@echo "Running atlas migrate down"
	@./scripts/atlas-apply.sh down

#apply migration
migrate-clean:
	@echo "Running atlas migrate clean"
	@./scripts/atlas-clean.sh clean

# Rollback migration state
migrate-diff-clean: 
	@echo "Running atlas migrate schema clean"
	@./scripts/atlas-clean.sh 


.PHONY: all build run test clean
