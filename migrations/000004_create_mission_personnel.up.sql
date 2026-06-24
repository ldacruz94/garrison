CREATE TABLE mission_personnel (
  id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  mission_id   UUID        NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
  personnel_id UUID        NOT NULL REFERENCES personnel(id) ON DELETE CASCADE,
  role         VARCHAR     NOT NULL,
  assigned_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
