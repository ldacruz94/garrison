# 007 - Graceful Shutdown

**Status:** Planned

## Scope

Handle `SIGINT` and `SIGTERM` so the server finishes in-flight requests before exiting instead of dropping them.

## Acceptance Criteria

- Server listens for `os.Signal` on a channel
- On signal, calls `server.Shutdown(ctx)` with a reasonable timeout (e.g. 10s)
- In-flight requests complete before the process exits
- DB connection pool is closed cleanly on shutdown
