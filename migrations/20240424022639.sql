-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "birth_date" TYPE timestamptz using birth_date::timestamptz;
