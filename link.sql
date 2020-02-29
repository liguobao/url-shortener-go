CREATE TABLE `link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(100) DEFAULT NULL,
  `short_code` varchar(30) DEFAULT NULL,
  `url` varchar(170) DEFAULT NULL,
  `update_time` datetime(6) DEFAULT NULL,
  `create_time` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url_user` (`open_id`,`url`),
  KEY `idx_code` (`short_code`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000000;
