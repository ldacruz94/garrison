.PHONY: up down build run run-local seed migrate-up migrate-down logs gen-ca gen-server-cert gen-client-cert gen-certs

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose up --build -d

run:
	go run cmd/server/main.go

run-local:
	DATABASE_URL=postgres://user:password@localhost:5432/garrison?sslmode=disable MTLS_ENABLED=false go run cmd/server/main.go

seed:
	docker compose exec -T db psql -U user -d garrison -f /dev/stdin < scripts/seed.sql

migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down

logs:
	docker compose logs -f api

gen-ca:
	mkdir -p certs
	openssl genrsa -out certs/ca.key 4096
	openssl req -new -x509 -days 3650 -key certs/ca.key -out certs/ca.crt -subj "/CN=GarrisonCA"

gen-server-cert:
	openssl genrsa -out certs/server.key 4096
	openssl req -new -key certs/server.key -out certs/server.csr -subj "/CN=localhost"
	openssl x509 -req -days 365 -in certs/server.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/server.crt

gen-client-cert:
	openssl genrsa -out certs/client.key 4096
	openssl req -new -key certs/client.key -out certs/client.csr -subj "/CN=garrison-client"
	openssl x509 -req -days 365 -in certs/client.csr -CA certs/ca.crt -CAkey certs/ca.key -CAcreateserial -out certs/client.crt

gen-certs: gen-ca gen-server-cert gen-client-cert
