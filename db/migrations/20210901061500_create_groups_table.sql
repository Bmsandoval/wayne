-- +goose Up
CREATE TABLE IF NOT EXISTS `groups` (
    id binary(16) PRIMARY KEY,
    sub varchar(36) generated always as
        (insert(
            insert(
                insert(
                    insert(hex(id),9,0,'-'),
                14,0,'-'),
            19,0,'-'),
        24,0,'-')
    ) virtual,
    name VARCHAR (100),
    owner_user_id binary(16) REFERENCES users ON DELETE CASCADE,
    created_at  DATETIME DEFAULT NOW(),
    updated_at  DATETIME DEFAULT NOW(),
    deleted_at  DATETIME
);
CREATE INDEX idx_groups_owner_user_id ON `groups`(owner_user_id);

CREATE TABLE IF NOT EXISTS users_groups (
    id SERIAL PRIMARY KEY,
    user_id binary(16) REFERENCES users ON DELETE CASCADE,
    group_id binary(16) REFERENCES `groups` ON DELETE CASCADE
);
CREATE INDEX idx_users_groups_user_id ON users_groups(user_id);
CREATE INDEX idx_users_groups_group_id ON users_groups(group_id);
-- +goose Down
DROP INDEX idx_users_groups_user_id ON users_groups;
DROP INDEX idx_users_groups_group_id ON users_groups;
DROP TABLE IF EXISTS users_groups;

DROP INDEX idx_groups_owner_user_id on `groups`;
DROP TABLE IF EXISTS `groups`;
