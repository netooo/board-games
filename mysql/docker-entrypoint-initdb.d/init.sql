CREATE DATABASE IF NOT EXISTS `boardgames` COLLATE 'utf8mb4_general_ci' ;
CREATE DATABASE IF NOT EXISTS `boardgames_test` COLLATE 'utf8mb4_general_ci' ;

GRANT ALL ON `board-games`.* TO 'user'@'%' ;
GRANT ALL ON `board-games_test`.* TO 'user'@'%' ;

FLUSH PRIVILEGES ;
