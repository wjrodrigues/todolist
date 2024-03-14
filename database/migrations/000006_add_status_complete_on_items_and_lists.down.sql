BEGIN;

ALTER TABLE items DROP CONSTRAINT items_check;
ALTER TABLE items
ADD CONSTRAINT items_check CHECK (
  (
    (status)::text = ANY (ARRAY['pending'::text, 'in_progress'::text, 'canceled'::text])
  )
);

ALTER TABLE lists DROP CONSTRAINT lists_check;
ALTER TABLE lists
ADD CONSTRAINT items_check CHECK (
  (
    (status)::text = ANY (ARRAY['pending'::text, 'in_progress'::text, 'canceled'::text])
  )
);

COMMIT;
