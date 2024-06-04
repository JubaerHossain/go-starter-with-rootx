# Variables
IMAGE_NAME := golang-restaurant-image
CONTAINER_NAME := golang-restaurant-container

# Targets
migrate-create:
    migrate create -ext sql -dir migrations -seq $(name)
install:
	go mod tidy
	swag init -g ./cmd/server/main.go
lint:
	golangci-lint run

seed:
	go run ./cmd/seed/seed.go

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

.PHONY: install seed dev web build run deploy docker-stop docker-remove docker-clean command cpu docs migrate-create
