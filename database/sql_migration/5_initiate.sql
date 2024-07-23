-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    audience_id INT NOT NULL,
    full_name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (audience_id) REFERENCES audiences(user_id) ON DELETE CASCADE
);

-- +migrate StatementEnd