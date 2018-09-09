-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_match` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'いいねを先に送ったユーザー',
  `partner_id` bigint(20) unsigned NOT NULL COMMENT 'ありがとういいねを送ったユーザー',
  `created_at` datetime NOT NULL COMMENT 'レコード作成日時',
  `updated_at` datetime NOT NULL COMMENT 'レコード更新日時',
  PRIMARY KEY (`user_id`, `partner_id`),
  KEY `idx_partner_id_user_id` (`partner_id`, `user_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='マッチ';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user_match`;
