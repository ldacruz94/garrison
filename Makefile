.PHONY: up down build run seed migrate-up migrate-down logs

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose up --build -d

run:
	go run cmd/server/main.go

seed:
	docker compose exec -T db psql -U user -d garrison -f /dev/stdin < scripts/seed.sql

migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down

logs:
	docker compose logs -f api
