CREATE TABLE IF NOT EXISTS email_text (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subject TEXT NOT NULL,
    body TEXT NOT NULL,
    status BOOLEAN
);

CREATE TABLE IF NOT EXISTS email_send_email (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    email VARCHAR(100) NOT NULL,
    send_time TIMESTAMP,
    send_status BOOLEAN NOT NULL DEFAULT FALSE,
    text_id UUID NOT NULL REFERENCES email_text(id)
);

CREATE TABLE IF NOT EXISTS tasks(
    Id text,
    Name text NOT NULL,
    Deadline timestamp NOT NULL,
    Summary text,
    AssigneId text NOT NULL,
    Status text NOT NULL,
    Create_at timestamp NOT NULL,
    Update_at timestamp,
    Delete_at timestamp

);

CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username text NOT NULL,
    profile_photo text,
    bio text,
    email text NOT NULL,
    gender text NOT NULL,
    PASSWORD text NOT NULL,
    adress jsonb NOT NULL,
    phone_numbers jsonb NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP ,
    deleted_at TIMESTAMP 
);

ALTER TABLE users ADD COLUMN
    refresh_taken text,
     ADD COLUMN acsess_token text;

CREATE UNIQUE INDEX IF NOT EXISTS user_index 
ON users(username) WHERE deleted_at IS NULL;

CREATE UNIQUE INDEX IF NOT EXISTS email_index 
ON users(email) WHERE deleted_at IS NULL;