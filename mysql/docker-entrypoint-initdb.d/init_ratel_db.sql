DROP TABLE IF EXISTS `soundcloud_track`;

CREATE TABLE `soundcloud_track` (
  `track_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `title` varchar(32) NOT NULL,
  `author` varchar(32) NOT NULL DEFAULT '',
  `description` varchar(1024) NOT NULL DEFAULT '',
  `display_time` datetime NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`track_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `soundcloud_track` WRITE;
/*!40000 ALTER TABLE `soundcloud_track` DISABLE KEYS */;

INSERT INTO `soundcloud_track` (`track_id`, `name`, `title`, `author`, `description`, `display_time`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(347152008,'etude_8','Etude 8','Hisanori ITo','','2017-10-16 00:00:00','2017-10-28 13:12:56','2017-10-23 00:00:00',NULL),
	(347625685,'blog_8','BLOG 8','yokushiryoku','','2017-10-20 00:00:00','2017-10-28 13:12:54','2017-10-23 00:00:00',NULL),
	(348067328,'etude_9','Etude 9','Hisanori ITo','','2017-10-24 00:00:00','2017-10-28 13:12:58','2017-10-23 00:00:00',NULL);

/*!40000 ALTER TABLE `soundcloud_track` ENABLE KEYS */;
UNLOCK TABLES;


/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
