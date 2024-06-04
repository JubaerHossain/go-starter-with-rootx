# Variables
IMAGE_NAME := go-starter-image
CONTAINER_NAME := go-starter-container

# Targets
migration-create:
	@echo "Creating migration file..."
	@echo "Migration name: $(name)"
	@echo "Table name: $(table)"
	@timestamp=$$(date +'%Y%m%d%H%M%S'); \
	filename="migrations/$$timestamp-$(name).sql"; \
	echo "-- Migration $(name)" > $$filename; \
	echo "" >> $$filename; \
	echo "CREATE TABLE IF NOT EXISTS $(name) (" >> $$filename; \
	echo "    id SERIAL PRIMARY KEY," >> $$filename; \
	echo "    name VARCHAR(100) NOT NULL," >> $$filename; \
	echo "    description TEXT NOT NULL," >> $$filename; \
	echo "    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP," >> $$filename; \
	echo "    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP" >> $$filename; \
	echo ");" >> $$filename; \
	echo "Migration file created: $$filename"

migration-up:
	@echo "Applying migrations..."
	@psql postgres://$(name):$(password)@$(host):$(port)/$(dbname) -f migrations/*.sql
	@echo "Migrations applied successfully"

install:
	go mod tidy
	swag init -g ./cmd/server/main.go
lint:
	golangci-lint run

seed:
	@echo "Running seeders..."
	@for file in seeds/*.sql; do \
		psql postgres://$(name):$(password)@$(host):$(port)/$(dbname) -f $$file; \
	done
	@echo "Seeders executed successfully"

docs:
	swag init -g ./cmd/server/main.go

dev:
	go run ./cmd/server

web:
	go run ./cmd/web/web.go

build:
	go build -o ./bin/server ./cmd/server

run:
	./bin/server

deploy:
	docker-compose -f docker-compose.yaml up -d

re-deploy:
	docker-compose -f docker-compose.yaml down
	docker system prune -f
	docker-compose -f docker-compose.yaml up -d --build

docker-stop:
	docker stop $(CONTAINER_NAME)

docker-remove:
	docker rm $(CONTAINER_NAME)

docker-clean:
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

cpu:
	echo "CPU Usage: "$[100-$(vmstat 1 2|tail -1|awk '{print $15}')]"%"

command:
	go run ./cmd/clid create github.com/JubaerHossain/rootx ${name}

.PHONY: install seed dev web build run deploy docker-stop docker-remove docker-clean command cpu docs migrate-create migrate-up
