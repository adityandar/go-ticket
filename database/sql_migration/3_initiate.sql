-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE audiences (
    user_id INT PRIMARY KEY,
    phone VARCHAR(256) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate StatementEnd