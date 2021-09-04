-- +goose Up
CREATE TABLE IF NOT EXISTS sessions (
    id serial PRIMARY KEY,
    name VARCHAR (50),
    user_id binary(16) REFERENCES users ON DELETE CASCADE,
    created_at  DATETIME DEFAULT NOW(),
    updated_at  DATETIME DEFAULT NOW(),
    deleted_at  DATETIME
);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
-- +goose Down
DROP INDEX idx_sessions_user_id on sessions;
DROP TABLE IF EXISTS sessions;
