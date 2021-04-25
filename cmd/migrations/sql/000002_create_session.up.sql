CREATE TABLE sessions (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `session_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'セッションID',
    `data` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'データ',
    `user_id` bigint(20) unsigned NOT NULL COMMENT 'ユーザID',
    `created_at` TIMESTAMP COMMENT '作成日時',
    `updated_at` TIMESTAMP COMMENT '更新日時',
    `deleted_at` TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY (`id`),
    KEY `foreign_sessions_on_user_id` (`user_id`), /* アホなので外部キー張れなかった。誰かお願い。 */
    UNIQUE KEY `uniq_sessions_on_session_id` (`session_id`),
    UNIQUE KEY `uniq_sessions_on_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;