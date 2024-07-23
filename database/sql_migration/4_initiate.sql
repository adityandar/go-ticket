-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    organizer_id INT NOT NULL,
    title VARCHAR(256) NOT NULL,
    date_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (organizer_id) REFERENCES organizers(user_id) ON DELETE CASCADE
);

-- +migrate StatementEnd