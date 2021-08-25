-- +goose Up
CREATE TABLE IF NOT EXISTS `greetings`
(
    `id`        INT          NOT NULL        AUTO_INCREMENT,
    `value`     VARCHAR(255) NOT NULL,

    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),

    INDEX `greetings_created_at_idx` (`created_at` ASC),
    INDEX `greetings_updated_at_idx` (`updated_at` ASC)
);
-- +goose Down
DROP TABLE IF EXISTS `greetings`;
