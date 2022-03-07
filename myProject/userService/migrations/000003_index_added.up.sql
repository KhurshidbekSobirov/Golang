CREATE UNIQUE INDEX IF NOT EXISTS user_index 
ON users(username) WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS email_index 
ON users(email) WHERE deleted_at IS NULL;