.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build up down

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down