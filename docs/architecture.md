# Architecture

## Overview

Go HTTP server backed by PostgreSQL. Clients connect over mTLS.

```
Client (cert required) <---mTLS---> Garrison API (Go) <---SQL---> PostgreSQL
```

## Auth

mTLS only. No API keys, no JWTs. Both sides present certs signed by the same CA and the connection is rejected at the handshake if either side can't verify the other.

## Data model

**Missions** are the core entity with a type, status, and optional time window.

**Personnel** are service members with a rank, unit, and clearance level, linked to missions through `MISSION_PERSONNEL`.

**Assets** are equipment or systems linked to missions through `MISSION_ASSET`.

**Audit log** is append-only. Every create, update, or delete writes a row with the before/after state in JSONB.

## Deployment

Docker Compose spins up the API and Postgres together. See the README for running locally with mTLS disabled.
