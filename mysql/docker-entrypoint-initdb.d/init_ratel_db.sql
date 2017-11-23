# ************************************************************
# Sequel Pro SQL dump
# バージョン 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# ホスト: 127.0.0.1 (MySQL 5.7.20)
# データベース: ratel
# 作成時刻: 2017-11-23 16:57:42 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# テーブルのダンプ entry
# ------------------------------------------------------------

DROP TABLE IF EXISTS `entry`;

CREATE TABLE `entry` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(32) NOT NULL DEFAULT '',
  `author` varchar(32) NOT NULL DEFAULT '',
  `body` varchar(4096) NOT NULL DEFAULT '',
  `display_at` datetime NOT NULL,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT NULL,
  `delete_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `entry` WRITE;
/*!40000 ALTER TABLE `entry` DISABLE KEYS */;

INSERT INTO `entry` (`id`, `title`, `author`, `body`, `display_at`, `create_at`, `update_at`, `delete_at`)
VALUES
	(1,'初ブログ！','hidelberq','初ブログです。\nサイトに新機能を色々つけてました。\n\n- ブログ機能 (今書いているここ)\n- お問合わせのフォーム\n- Music Log のデータをデータベースで管理する\n\n特に3番めが重要です。\n今まで Ito くんが音楽をつくってくれてるのにもかかわらず、\nHTMLを直接編集するのが辛くて更新してなかったですが、(すみません。。)\nこれからは、手軽に更新できるようになりました。\n\nRatel の新鮮な情報を発信して行きたいとおもうので、\nこれからもアクセスよろしくお願いします。。！','2017-11-23 15:13:00','2017-11-23 16:54:04',NULL,NULL);

/*!40000 ALTER TABLE `entry` ENABLE KEYS */;
UNLOCK TABLES;


# テーブルのダンプ message
# ------------------------------------------------------------

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `email` varchar(128) DEFAULT '',
  `body` varchar(4096) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# テーブルのダンプ soundcloud_track
# ------------------------------------------------------------

DROP TABLE IF EXISTS `soundcloud_track`;

CREATE TABLE `soundcloud_track` (
  `track_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `title` varchar(32) NOT NULL DEFAULT '',
  `author` varchar(32) NOT NULL DEFAULT '',
  `description` varchar(1024) NOT NULL DEFAULT '',
  `display_time` datetime NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`track_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `soundcloud_track` WRITE;
/*!40000 ALTER TABLE `soundcloud_track` DISABLE KEYS */;

INSERT INTO `soundcloud_track` (`track_id`, `name`, `title`, `author`, `description`, `display_time`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(347152008,'etude_8','Etude 8','Hisanori ITo','','2017-10-16 00:00:00','2017-10-28 13:12:56','2017-10-23 00:00:00',NULL),
	(347625685,'blog_8','BLOG 8','yokushiryoku','','2017-10-20 00:00:00','2017-10-28 13:12:54','2017-10-23 00:00:00',NULL),
	(348067328,'etude_9','Etude 9','Hisanori ITo','','2017-10-24 00:00:00','2017-10-28 13:12:58','2017-10-23 00:00:00',NULL);

/*!40000 ALTER TABLE `soundcloud_track` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
