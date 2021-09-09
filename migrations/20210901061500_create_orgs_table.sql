-- +goose Up
CREATE TABLE IF NOT EXISTS orgs (
    id binary(16) PRIMARY KEY,
    sub varchar(36) generated always as (
        insert(
            insert(
                insert(
                    insert( hex(id),9,0,'-' ),
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

CREATE INDEX idx_orgs_owner_user_id ON orgs(owner_user_id);

ALTER TABLE users ADD org_id binary(16) REFERENCES orgs ON DELETE CASCADE;
CREATE INDEX idx_users_org_id ON users(org_id);
-- +goose Down
DROP INDEX idx_users_org_id ON users;
ALTER TABLE users DROP org_id;

DROP INDEX idx_orgs_owner_user_id on orgs;
DROP TABLE IF EXISTS orgs;
