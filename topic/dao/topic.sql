CREATE TABLE `topic` (
  `tid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '@desc pk',
  `uid` int(11) DEFAULT NULL COMMENT '@desc uid',
  `title` varchar(200) DEFAULT NULL COMMENT '@desc title',
  `content` varchar(255) DEFAULT NULL COMMENT '@desc content',
  `like` tinyint(255) DEFAULT NULL COMMENT '@desc the count of like',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '@desc the time publish',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '@desc the time update',
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `comment` (
  `id` int(11) NOT NULL COMMENT '@desc pk',
  `content` varchar(255) DEFAULT NULL COMMENT '@desc commont content',
  `uid` int(11) DEFAULT NULL COMMENT '@desc uid',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '@desc ',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '@desc ',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;