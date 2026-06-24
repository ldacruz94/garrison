# Architecture

## Overview

Simple setup — Go HTTP server talking to PostgreSQL. Clients connect over mTLS, so there's no session layer or token management to worry about.

```
Client (cert required) <---mTLS---> Garrison API (Go) <---SQL---> PostgreSQL
```

## Auth

mTLS only. No API keys, no JWTs. The client and server both present certs signed by the same CA, and the connection is rejected at the TLS handshake if either side can't verify the other. The authenticated client identity gets written as `actor_id` on every audit log entry.

## Data model

**Missions** are the core entity. Each one has a type (recon, logistics, etc.), a status, and an optional start/end window.

**Personnel** are service members with a rank, unit, and clearance level. They're linked to missions via `MISSION_PERSONNEL`, which captures what role they're filling on that mission.

**Assets** are physical or virtual resources — vehicles, radios, systems. Same pattern as personnel: linked to missions via `MISSION_ASSET` with a role attached.

**Audit log** is append-only. Every create, update, or delete writes a row with the entity that changed, who changed it, and full before/after snapshots in JSONB. Nothing in this table ever gets modified or deleted.

## API

Standard REST, JSON bodies, UUID primary keys.

| Method | Route | What it does |
|---|---|---|
| GET | `/missions` | List missions |
| POST | `/missions` | Create a mission |
| GET | `/missions/{id}` | Get a mission |
| PUT | `/missions/{id}` | Update a mission |
| DELETE | `/missions/{id}` | Delete a mission |
| POST | `/missions/{id}/personnel` | Assign personnel |
| POST | `/missions/{id}/assets` | Assign an asset |
| GET | `/personnel` | List personnel |
| GET | `/assets` | List assets |
| GET | `/audit-log` | Query audit history |

## Deployment

Docker Compose for local dev — spins up the API server and Postgres together. In a real environment this sits behind a network boundary where only clients with valid certs can reach it.
