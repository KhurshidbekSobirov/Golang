CREATE TABLE users(
    id UUID NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username text NOT NULL,
    profile_photo text,
    bio text,
    email text NOT NULL,
    gender text NOT NULL,
    adress jsonb NOT NULL,
    phone_numbers jsonb NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP ,
    deleted_at TIMESTAMP 
);