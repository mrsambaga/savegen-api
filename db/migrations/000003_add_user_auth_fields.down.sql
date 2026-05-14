DROP INDEX IF EXISTS users_google_sub_unique;
DROP INDEX IF EXISTS users_email_unique;

ALTER TABLE users
    DROP COLUMN IF EXISTS google_sub,
    DROP COLUMN IF EXISTS is_guest,
    DROP COLUMN IF EXISTS password_hash;
