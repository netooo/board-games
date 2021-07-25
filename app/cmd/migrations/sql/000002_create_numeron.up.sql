CREATE TABLE numerons (
    `id` integer unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `status` tinyint(1) unsigned NOT NULL COMMENT 'ステータス',
    `owner_id` bigint(20) unsigned NOT NULL COMMENT '作成者プレイヤーID',
    `created_at` TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;