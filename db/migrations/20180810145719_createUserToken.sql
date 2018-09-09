-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_token` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
  `token` char(40) NOT NULL COMMENT 'アクセストークン',
  `created_at` datetime NOT NULL COMMENT 'レコード作成日時',
  `updated_at` datetime NOT NULL COMMENT 'レコード更新日時',
  PRIMARY KEY (`user_id`),
  KEY `idx_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='トークン';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user_token`;
