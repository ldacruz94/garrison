
CREATE TABLE personnel (
  id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  rank            VARCHAR NOT NULL,
  last_name       VARCHAR NOT NULL,
  first_name      VARCHAR NOT NULL,
  unit_designator VARCHAR NOT NULL,
  clearance_level VARCHAR NOT NULL,
  status          VARCHAR NOT NULL,
  created_at      TIMESTAMPTZ DEFAULT NOW(),
  updated_at      TIMESTAMPTZ DEFAULT NOW()
);
