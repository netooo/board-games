CREATE DATABASE IF NOT EXISTS `boardgames` COLLATE 'utf8mb4_general_ci' ;
CREATE DATABASE IF NOT EXISTS `boardgames_test` COLLATE 'utf8mb4_general_ci' ;

GRANT ALL ON `boardgames`.* TO 'user'@'%' ;
GRANT ALL ON `boardgames_test`.* TO 'ser'@'%' ;

FLUSH PRIVILEGES ;
