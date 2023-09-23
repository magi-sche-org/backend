-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE `event`
ADD COLUMN `enables_email_notification` BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN `expected_participants_number` INT NOT NULL DEFAULT 0,
    ADD COLUMN `notification_email` VARCHAR(255) NOT NULL DEFAULT '';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE `event` DROP COLUMN `enables_email_notification`,
    DROP COLUMN `expected_participants_number`,
    DROP COLUMN `notification_email`;
-- +goose StatementEnd
