-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user` (
    `id` CHAR(26) NOT NULL,
    -- set as default nicknameとかあとでやっても良いかも
    `name` VARCHAR(128) NOT NULL,
    `is_registered` BOOLEAN NOT NULL DEFAULT FALSE,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`)
);
CREATE TABLE `event` (
    `id` CHAR(26) NOT NULL,
    `owner_id` CHAR(26) NOT NULL,
    `name` VARCHAR(128) NOT NULL,
    `description` TEXT NOT NULL,
    `duration_about` VARCHAR(128) NOT NULL,
    `unit_seconds` BIGINT UNSIGNED NOT NULL,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_events_owner_id` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);
CREATE TABLE `event_time_unit` (
    `id` CHAR(26) NOT NULL,
    `event_id` CHAR(26) NOT NULL,
    `time_slot` DATETIME NOT NULL,
    `slot_seconds` BIGINT UNSIGNED NOT NULL,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_unit_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `event` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE `user_event_answer` (
    `id` CHAR(26) NOT NULL,
    `user_id` CHAR(26) NOT NULL,
    `event_id` CHAR(26) NOT NULL,
    `user_nickname` VARCHAR(128) NOT NULL,
    `note` TEXT NOT NULL,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_answers_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_answers_event_id` FOREIGN KEY (`event_id`) REFERENCES `event` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_answers_unique` UNIQUE (`user_id`, `event_id`)
);
CREATE TABLE `user_event_answer_unit` (
    `id` CHAR(26) NOT NULL,
    `user_event_answer_id` CHAR(26) NOT NULL,
    `event_time_unit_id` CHAR(26) NOT NULL,
    `availability` ENUM('available', 'maybe', 'unavailable') NOT NULL,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_unit_status_user_event_answer_id` FOREIGN KEY (`user_event_answer_id`) REFERENCES `user_event_answer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_unit_status_event_time_unit_id` FOREIGN KEY (`event_time_unit_id`) REFERENCES `event_time_unit` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_unit_status_unique` UNIQUE (`user_event_answer_id`, `event_time_unit_id`)
);
CREATE TABLE `refresh_token` (
    `id` CHAR(26) NOT NULL,
    `user_id` CHAR(26) NOT NULL,
    `token` CHAR(64) NOT NULL,
    `expires_at` DATETIME(6) NOT NULL,
    `revoked` BOOLEAN NOT NULL DEFAULT FALSE,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_token_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    INDEX `idx_token` (`token`)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `refresh_token`;
DROP TABLE IF EXISTS `user_event_answer_unit`;
DROP TABLE IF EXISTS `user_event_answer`;
DROP TABLE IF EXISTS `event_time_unit`;
DROP TABLE IF EXISTS `event`;
DROP TABLE IF EXISTS `user`;
-- +goose StatementEnd
