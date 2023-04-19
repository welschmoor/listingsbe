CREATE TABLE IF NOT EXISTS listings (
  id bigserial PRIMARY KEY,
  price integer NOT NULL,
  title text NOT NULL,
  description text NOT NULL,
  categories text [] NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  version integer NOT NULL DEFAULT 1
);