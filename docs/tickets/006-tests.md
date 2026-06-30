# 006 - Tests

**Status:** Planned  
**Depends on:** 005 (interface-driven stores)

## Scope

Add tests at two levels: unit tests for handlers using mock stores, and integration tests that hit a real DB.

## Acceptance Criteria

- Handler tests use mock stores (no DB required)
- Integration tests run against a real PostgreSQL instance
- Table-driven test style throughout
- At minimum, cover CRUD for Missions and the middleware behaviors from ticket 001
