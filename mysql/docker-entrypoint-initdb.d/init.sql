CREATE DATABASE IF NOT EXISTS `board-games` COLLATE 'utf8mb4_general_ci' ;
CREATE DATABASE IF NOT EXISTS `board-games_test` COLLATE 'utf8mb4_general_ci' ;

GRANT ALL ON `board-games`.* TO 'user'@'%' ;
GRANT ALL ON `board-games_test`.* TO 'user'@'%' ;

FLUSH PRIVILEGES ;
