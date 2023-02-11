CREATE TABLE `users` (
    `id` VARCHAR(36) NOT NULL COMMENT 'ユーザーID',
    `name` VARCHAR(128) NOT NULL COMMENT 'ユーザー名',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`)
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'ユーザー';
CREATE TABLE `events` (
    `id` VARCHAR(36) NOT NULL COMMENT 'イベントID',
    `owner_id` VARCHAR(36) NOT NULL COMMENT 'イベントオーナーID',
    `name` VARCHAR(128) NOT NULL COMMENT 'イベントのタイトル',
    `description` TEXT COMMENT 'イベントの説明',
    `duration` BIGINT UNSIGNED NOT NULL COMMENT 'イベントの長さ',
    `unit_second` BIGINT UNSIGNED NOT NULL COMMENT '時間単位',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_events_owner_id` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'タスク';
CREATE TABLE `unit_events` (
    `id` VARCHAR(36) NOT NULL COMMENT 'スケジュールID',
    `event_id` VARCHAR(36) NOT NULL COMMENT 'イベントID',
    `start_at` DATETIME NOT NULL COMMENT 'スケジュールの開始日時',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_unit_events_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'イベント単位';
CREATE TABLE `answers` (
    `id` VARCHAR(36) NOT NULL COMMENT '回答ID',
    `user_id` VARCHAR(36) NOT NULL COMMENT '回答者ID',
    `event_id` VARCHAR(36) NOT NULL COMMENT 'イベントID',
    `name` VARCHAR(128) NOT NULL COMMENT '回答者の名前',
    `note` TEXT COMMENT '回答者の備考',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_answers_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `fk_answers_event_id` FOREIGN KEY (`event_id`) REFERENCES `events` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '回答';
CREATE TABLE `unit_statuses` (
    `id` VARCHAR(36) NOT NULL COMMENT 'ステータスID',
    `answer_id` VARCHAR(36) NOT NULL COMMENT '回答ID',
    `start_at` DATETIME NOT NULL COMMENT 'スケジュールの開始日時',
    `status` INT UNSIGNED NOT NULL COMMENT 'ステータス',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_unit_statuses_answer_id` FOREIGN KEY (`answer_id`) REFERENCES `answers` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'ステータス';
CREATE TABLE `tokens` (
    `id` VARCHAR(36) NOT NULL COMMENT 'トークンID',
    `user_id` VARCHAR(36) NOT NULL COMMENT 'ユーザーID',
    `created_at` DATETIME(6) NOT NULL COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL COMMENT 'レコード修正日時',
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_token_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'トークン';
