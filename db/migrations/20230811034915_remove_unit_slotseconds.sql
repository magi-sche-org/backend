-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE `event_time_unit` DROP COLUMN `slot_seconds`;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE `event_time_unit`
ADD COLUMN `slot_seconds` BIGINT UNSIGNED NOT NULL DEFAULT 0;
-- eventのunit_secondsの値を，これにぶち込む（event||--|{event_time_unit）
UPDATE `event_time_unit` AS etu
    JOIN `event` AS e ON etu.event_id = e.id
SET etu.slot_seconds = e.unit_seconds;
-- +goose StatementEnd
