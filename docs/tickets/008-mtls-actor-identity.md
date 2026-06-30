# 008 - mTLS Client Identity in Audit Log

**Status:** Planned  
**Depends on:** 001 (middleware)

## Scope

Extract the client identity from the mTLS certificate on each request and wire it into the audit log as `actor_id`. Right now `actor_id` is unused.

## Acceptance Criteria

- Middleware extracts the client CN from `r.TLS.PeerCertificates[0]`
- Identity is attached to the request context
- Handlers (or store layer) pull the identity from context and write it as `actor_id` on every audit log entry
- Requests with no peer certificate (mTLS disabled mode) set `actor_id` to a sentinel value or null
