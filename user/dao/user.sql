CREATE TABLE `user` (
  `uid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '@desc pk',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '@desc user_name',
  `age` int(11) NOT NULL COMMENT '@desc user age',
  `phone` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '@desc phone',
  `sex` tinyint(4) NOT NULL COMMENT '@desc sex @value 1-male 2-female 3-no sex',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '@desc create_time',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '@desc update_time',
  `yunxin_token` varchar(255) DEFAULT NULL COMMENT '@desc yunxin',
  `yunxin_id` varchar(255) DEFAULT NULL COMMENT '@desc yunxin id',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `index_uniq_user_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;