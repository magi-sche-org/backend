-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE `oauth_provider` (
    `id` CHAR(26) NOT NULL,
    `name` VARCHAR(128) NOT NULL,
    `client_id` VARCHAR(128) NOT NULL,
    `client_secret` VARCHAR(128) NOT NULL,
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_oauth_provider_unique` UNIQUE (`name`)
);
CREATE TABLE `oauth_user_info` (
    `id` CHAR(26) NOT NULL,
    `user_id` CHAR(26) NOT NULL,
    `provider_id` CHAR(26) NOT NULL,
    `provider_uid` VARCHAR(128) NOT NULL,
    `access_token` VARCHAR(2048) NOT NULL,
    `refresh_token` VARCHAR(512) NOT NULL,
    `access_token_expires_at` DATETIME(6) NOT NULL,
    `refresh_token_expires_at` DATETIME(6),
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_oauth_user_info_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_oauth_user_info_provider_id` FOREIGN KEY (`provider_id`) REFERENCES `oauth_provider` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_oauth_user_info_provider_unique` UNIQUE (`user_id`, `provider_id`),
    CONSTRAINT `fk_oauth_user_info_unique` UNIQUE (`provider_id`, `provider_uid`)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE `oauth_user_info`;
DROP TABLE `oauth_provider`;
-- +goose StatementEnd
