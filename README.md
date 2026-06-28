# Garrison

REST API for managing DoD missions, personnel, and assets. Built in Go with PostgreSQL.

## What it does

Garrison is really just a simple CRUD-based API that lets you create and track missions, assign service members and equipment to them, and keeps a full audit log of every change. All traffic is secured with mutual TLS so both sides of every connection have to prove who they are.

## Entities

| Entity | What it is |
|---|---|
| **Mission** | An operation with a type, status, and time window |
| **Personnel** | Service members — rank, unit, clearance level |
| **Asset** | Equipment or systems assigned to missions |
| **Audit Log** | Append-only record of every change across the system |

Personnel and assets are linked to missions through join tables that track role and assignment time. See [`docs/erd.md`](docs/erd.md) for the full schema.

NOTE: This is currently a WIP and I'll be adding additional entities as I mature out the project.

## Stack

- Go 1.22
- PostgreSQL
- mTLS
- Docker / Docker Compose

## Running it

```bash
# start the DB and API
make build

# stop everything
make down

# run locally (DB must be up)
make up
make run
```

### Seeding the database

```bash
make seed
```

Truncates and re-inserts sample missions, personnel, and assets. Safe to run multiple times.

## Security

For this type of project, I wanted to explore mTLS to better understand implementing it usin Golang.

