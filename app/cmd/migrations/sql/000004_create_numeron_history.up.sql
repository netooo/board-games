CREATE TABLE numeron_histories (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `numeron_id` integer unsigned NOT NULL COMMENT 'ヌメロンルームID',
    `player_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
    `action` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '回答番号',
    `result` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '結果',
    `turn` tinyint unsigned NOT NULL COMMENT 'ターン',
    `created_at` TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;