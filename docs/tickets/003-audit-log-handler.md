# 003 - Audit Log Handler

**Status:** Planned

## Scope

Wire up an HTTP handler for the audit log. The store and model already exist, this is adding the handler and routes.

## Endpoints

- `GET /audit-log` - list audit log entries, filterable by `entity_type` and `entity_id` via query params

## Acceptance Criteria

- Returns a paginated list of audit log entries
- Supports `?entity_type=mission` and `?entity_id=<uuid>` query params for filtering
- Returns entries in descending order by `occurred_at`
- No write endpoints, the audit log is append-only
