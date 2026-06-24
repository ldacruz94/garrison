
CREATE TABLE missions (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name          VARCHAR NOT NULL,
  description   TEXT,
  status        VARCHAR NOT NULL,
  mission_type  VARCHAR NOT NULL,
  start_time    TIMESTAMPTZ,
  end_time      TIMESTAMPTZ,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

