BEGIN;

ALTER TABLE items DROP CONSTRAINT items_check;
ALTER TABLE items
ADD CONSTRAINT items_check CHECK (
  (
    (status)::text = ANY (ARRAY['pending'::text, 'in_progress'::text, 'canceled'::text, 'completed'::text])
  )
);

ALTER TABLE lists DROP CONSTRAINT items_check;
ALTER TABLE lists
ADD CONSTRAINT lists_check CHECK (
  (
    (status)::text = ANY (ARRAY['pending'::text, 'in_progress'::text, 'canceled'::text, 'completed'::text])
  )
);

COMMIT;
