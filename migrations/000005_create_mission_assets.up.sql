CREATE TABLE mission_assets (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  mission_id  UUID        NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
  asset_id    UUID        NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
  role        VARCHAR     NOT NULL,
  assigned_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
