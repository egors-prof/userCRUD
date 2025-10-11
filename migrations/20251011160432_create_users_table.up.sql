create table if not exists users(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    hash_pass VARCHAR NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT TIMESTAMP,
    role VARCHAR default 'USER'
);