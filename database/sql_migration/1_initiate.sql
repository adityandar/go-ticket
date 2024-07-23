-- +migrate Up
-- +migrateStatementBegin

CREATE TYPE user_role AS ENUM ('admin', 'organizer', 'audience');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(256) UNIQUE NOT NULL,
    full_name VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +migrate StatementEnd