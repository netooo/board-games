CREATE TABLE `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID'
    `user_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'ユーザID'
    `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '氏名'
    `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'メールアドレス'
    `password` varchar(255) COLLATE utf8_inicode_ci NOT NULL COMMENT 'パスワード'
)