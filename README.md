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

NOTE: This is currently a WIP and I'll be adding additional entities and features as I mature out the project.

## Stack

- Go 1.25
- PostgreSQL 16
- mTLS (mutual TLS)
- Docker / Docker Compose

## Running it

### With Docker (mTLS enabled)

Generate certs first — only needed once:

```bash
make gen-certs
```

Then start everything:

```bash
make build
```

The API runs on `:8443` with mTLS enforced. Therefore, all clients are required to have a valid certificate signed by the generated CA.

```bash
# stop everything
make down
```

### Locally (mTLS disabled)

Start just the database, then run the API directly:

```bash
docker compose up -d db
make run-local
```

The API runs on `:8080` with no TLS. Use this if you're trying to run it locally and wanting to try out
the endpoints more easily on bruno/postman.

### Seeding the database

```bash
make seed
```

Truncates and re-inserts sample missions, personnel, and assets. Safe to run multiple times.

## mTLS

All production traffic is secured with mutual TLS. Both the client and server must present a certificate signed by the same CA (`GarrisonCA`). The generated certs live in `certs/` and are gitignored.

To test with curl against the Docker setup:

```bash
# Should succeed
curl --cacert certs/ca.crt \
     --cert certs/client.crt \
     --key certs/client.key \
     https://localhost:8443/missions

# Should fail — no client cert
curl --cacert certs/ca.crt https://localhost:8443/missions
```

