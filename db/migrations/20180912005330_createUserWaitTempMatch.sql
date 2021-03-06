-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_wait_temp_match` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'リクエストしたユーザID',
  `gender` enum('M','F') NOT NULL COMMENT '性別',
  `is_matched` boolean NOT NULL COMMENT 'マッチ済みかどうか',
  `is_canceled` boolean NOT NULL COMMENT 'リクエストキャンセル',
  `created_at` datetime NOT NULL COMMENT 'レコード作成日時',
  `updated_at` datetime NOT NULL COMMENT 'レコード更新日時',
  PRIMARY KEY (`user_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='マッチ待ち';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user_wait_temp_match`;
