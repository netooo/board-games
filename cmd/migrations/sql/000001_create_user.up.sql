CREATE TABLE users (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'ユーザID',
    `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '氏名',
    `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'メールアドレス',
    `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'パスワード',
    `created_at` TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_users_on_user_id` (`user_id`),
    UNIQUE KEY `uniq_users_on_email` (`email`)
)