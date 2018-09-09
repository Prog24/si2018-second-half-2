-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user_like` (
  `user_id` bigint(20) unsigned NOT NULL COMMENT 'いいねを送信したユーザー',
  `partner_id` bigint(20) unsigned NOT NULL COMMENT 'いいねをもらったユーザー',
  `created_at` datetime NOT NULL COMMENT 'レコード作成日時',
  `updated_at` datetime NOT NULL COMMENT 'レコード更新日時',
  PRIMARY KEY (`user_id`, `partner_id`, `created_at`),
  KEY `idx_partner_id_user_id` (`partner_id`, `user_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='いいね';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user_like`;
