-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `gender` enum('M','F') NOT NULL COMMENT '性別',
  `birthday` date NOT NULL COMMENT '誕生日',
  `nickname` varchar(20) NOT NULL COMMENT 'ニックネーム',
  `tweet` varchar(140) NOT NULL COMMENT 'つぶやき',
  `introduction` text NOT NULL COMMENT '自己紹介文',
  `residence_state` varchar(20) NOT NULL COMMENT '居住地',
  `home_state` varchar(20) NOT NULL COMMENT '出身地',
  `education` varchar(20) NOT NULL COMMENT '学歴',
  `job` varchar(20) NOT NULL COMMENT '職業',
  `annual_income` varchar(20) NOT NULL COMMENT '年収',
  `height` varchar(20) NOT NULL COMMENT '身長',
  `body_build` varchar(20) NOT NULL COMMENT '体型',
  `marital_status` varchar(20) NOT NULL COMMENT '結婚歴',
  `child` varchar(20) NOT NULL COMMENT '子供の有無',
  `when_marry` varchar(20) NOT NULL COMMENT '結婚に対する意思',
  `want_child` varchar(20) NOT NULL COMMENT '子供が欲しいか',
  `smoking` varchar(20) NOT NULL COMMENT 'タバコ',
  `drinking` varchar(20) NOT NULL COMMENT 'お酒',
  `holiday` varchar(20) NOT NULL COMMENT '休日',
  `how_to_meet` varchar(20) NOT NULL COMMENT '出会うまでの希望',
  `cost_of_date` varchar(20) NOT NULL COMMENT '初回デート費用',
  `nth_child` varchar(20) NOT NULL COMMENT '兄弟姉妹',
  `housework` varchar(20) NOT NULL COMMENT '家事・育児',
  `created_at` datetime NOT NULL COMMENT 'レコード作成日時',
  `updated_at` datetime NOT NULL COMMENT 'レコード更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `user`;
