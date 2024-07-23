-- +migrate Up
-- +migrateStatementBegin

CREATE TABLE organizers (
    user_id INT PRIMARY KEY,
    company_name VARCHAR(256) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate StatementEnd