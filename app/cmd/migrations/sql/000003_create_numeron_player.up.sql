CREATE TABLE numeron_players (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `numeron_id` integer unsigned NOT NULL COMMENT 'ヌメロンルームID',
    `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
    `order` tinyint unsigned NOT NULL COMMENT '順番',
    `code` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '設定番号',
    `rank` tinyint unsigned NOT NULL COMMENT '順位',
    `created_at` TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;