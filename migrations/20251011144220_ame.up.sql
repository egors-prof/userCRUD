create table users(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    hash_pass VARCHAR NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    role VARCHAR
);