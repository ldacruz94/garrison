CREATE TABLE audit_log (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  entity_type VARCHAR     NOT NULL,
  entity_id   UUID        NOT NULL,
  actor_id    UUID        NOT NULL REFERENCES personnel(id),
  action      VARCHAR     NOT NULL,
  old_value   JSONB,
  new_value   JSONB,
  occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
