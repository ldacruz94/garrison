# 001 - Middleware

**Status:** In Progress

## Scope

Add a `internal/middleware/` package with three middleware functions chained onto the mux in `main.go`.

- **Request Logger** - log method, path, status code, and duration per request
- **Panic Recovery** - catch unhandled panics, return 500, log the stack trace
- **Request Timeout** - wrap each request in a `context.WithTimeout` to prevent hung connections

## Acceptance Criteria

- All three middleware are in separate files under `internal/middleware/`
- Each follows the standard `func(next http.Handler) http.Handler` signature
- Middleware is chained in `main.go` wrapping the mux
- A request to any route logs method, path, status, and duration
- A panic in a handler returns a 500 and does not crash the server
- A request that exceeds the timeout returns a 503
