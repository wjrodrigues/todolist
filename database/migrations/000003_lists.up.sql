CREATE TABLE IF NOT EXISTS lists (
  id VARCHAR(255) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(500) NULL,
  status VARCHAR(50) NOT NULL,
  created_at TIME DEFAULT NOW(),
  updated_at TIME DEFAULT NOW(),
  CONSTRAINT items_check CHECK (((status)::text = ANY (ARRAY['pending'::text, 'in_progress'::text, 'canceled'::text])))
);
