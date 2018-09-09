-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_message` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'メッセージを送ったユーザー',
  `partner_id` bigint(20) unsigned NOT NULL COMMENT 'メッセージを受け取ったユーザー',
  `message` text COMMENT 'メッセージ本文',
  `created_at` datetime NOT NULL COMMENT 'レコード作成日時',
  `updated_at` datetime NOT NULL COMMENT 'レコード更新日時',
  PRIMARY KEY (`user_id`,`partner_id`,`created_at`),
  KEY `idx_message_user_partner` (`partner_id`,`user_id`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='メッセージ';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user_message`;
