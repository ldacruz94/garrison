# Roadmap

## In Progress

- Middleware (request logger, panic recovery, request timeout)

## Planned

- Join table endpoints (assign personnel and assets to missions)
- Audit log handler
- Clearance level enforcement
- Interface-driven store layer (for testability)
- Tests (table-driven, integration against real DB)
- Graceful shutdown
- Extract mTLS client identity from cert and wire into audit log `actor_id`

## Done

- CRUD for Missions, Personnel, Assets
- PostgreSQL with pgx and connection pooling
- Migrations and seed script
- mTLS with environment-based toggle for local dev
- Dockerized with Docker Compose
- Cert generation via Makefile
