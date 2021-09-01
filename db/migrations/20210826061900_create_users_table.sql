-- +goose Up
CREATE TABLE IF NOT EXISTS `users` (
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
   username VARCHAR(75),
   UNIQUE KEY unique_username (username),
   password binary(16),
   created_at  DATETIME DEFAULT NOW(),
   updated_at  DATETIME DEFAULT NOW(),
   deleted_at  DATETIME
);
-- +goose Down
DROP TABLE IF EXISTS users
