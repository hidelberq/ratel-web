# ************************************************************
# Sequel Pro SQL dump
# バージョン 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# ホスト: 104.155.194.243 (MySQL 5.7.20)
# データベース: ratel
# 作成時刻: 2017-12-06 15:13:37 +0000
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `entry` WRITE;
/*!40000 ALTER TABLE `entry` DISABLE KEYS */;

INSERT INTO `entry` (`id`, `title`, `author`, `body`, `display_at`, `create_at`, `update_at`, `delete_at`)
VALUES
	(1,'初ブログ。','hidelberq','初ブログです。\nサイトに新機能を色々つけてました。\n\n- ブログ機能 (今書いているここ)\n- お問合わせのフォーム\n- Music Log をデータベースで管理する\n\n特に3番めが重要です。\n今まで Ito くんが音楽をつくってくれてるの にもかかわらず、\nHTMLを直接編集するのが辛くて更新してなかったですが、(すみません。。)\nこれからは、手軽に更新できるようになりました。\n\nRatel の新鮮な情報を発信して行きたいとおもうので、\nこれからもアクセスよろしくお願いします。。！','2017-11-23 15:13:00','2017-11-23 17:32:53',NULL,NULL),
	(2,'a','a','a','2017-11-01 00:01:00','2017-11-27 15:40:53',NULL,'2017-11-28 00:40:54'),
	(3,'あ','あ','あ','2017-11-28 00:41:00','2017-11-28 16:43:45',NULL,'2017-11-29 01:43:46'),
	(4,'sexマシンガン2','sexマシンガン','sexマシンガン','2017-11-28 00:45:00','2017-11-28 16:43:09',NULL,'2017-11-29 01:43:09'),
	(5,'テスト(上書き用)','テスト','てすと','2017-11-28 00:00:00','2017-11-27 16:39:19',NULL,'2017-11-28 01:39:20'),
	(6,'あ','あ','い','2017-11-28 13:36:00','2017-11-28 16:42:51',NULL,'2017-11-29 01:42:52'),
	(7,'うんこ','うんこ','うんこ','2017-11-28 13:41:00','2017-11-28 16:27:32',NULL,'2017-11-29 01:27:33'),
	(8,'レーベルメイト募集中','Hisanori Ito','ビートメイカーのHisanoriです。\r\n\r\nRatelは近々自主レーベルを設立します。\r\n\r\n音楽じゃなくてもいいです。\r\n\r\nクリエイター募集中。','2017-11-28 14:17:00','2017-11-28 16:53:22',NULL,NULL),
	(9,'?','aaa?','aaa?','2017-11-28 16:27:00','2017-12-02 08:31:32',NULL,'2017-12-02 17:31:32');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;

INSERT INTO `message` (`id`, `name`, `email`, `body`, `created_at`)
VALUES
	(1,'a','a@a','','2017-11-28 16:51:20'),
	(2,'','','','2017-12-01 04:02:42'),
	(3,'','','','2017-12-01 04:02:54'),
	(4,'','','','2017-12-01 04:03:01'),
	(5,'','','','2017-12-03 22:42:58');

/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;


# テーブルのダンプ soundcloud_track
# ------------------------------------------------------------

DROP TABLE IF EXISTS `soundcloud_track`;

CREATE TABLE `soundcloud_track` (
  `track_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `title` varchar(64) NOT NULL DEFAULT '',
  `author` varchar(64) NOT NULL DEFAULT '',
  `description` varchar(1024) NOT NULL DEFAULT '',
  `display_at` datetime NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`track_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `soundcloud_track` WRITE;
/*!40000 ALTER TABLE `soundcloud_track` DISABLE KEYS */;

INSERT INTO `soundcloud_track` (`track_id`, `name`, `title`, `author`, `description`, `display_at`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'main','a','Hisanori Ito','','2017-11-27 00:00:00','2017-11-26 10:54:23','2017-11-26 19:54:23','2017-11-26 19:54:23'),
	(13,'a','aa','aaaa','','2017-11-29 00:00:00','2017-11-28 16:47:54','2017-11-29 01:47:55','2017-11-29 01:47:55'),
	(111,'a','abc','a','','2017-11-08 00:00:00','2017-11-28 16:47:45','2017-11-29 01:47:46','2017-11-29 01:47:46'),
	(1230,'125','a','a','','2017-11-26 00:00:00','2017-11-26 10:54:53','2017-11-26 19:54:53','2017-11-26 19:54:53'),
	(347152008,'etude_8','Etude 8','Hisanori Ito','','2017-10-16 00:00:00','2017-11-23 17:28:33','2017-10-23 00:00:00',NULL),
	(347625685,'blog_8','BLOG 8','抑止力 (track by Etude 8 Hisanori Ito)','来年のことなんか俺に聞くなっつんだよな！\r\nハイデルは世直ししたいらしいよ\r\nみんなこいつには投票するなよ','2017-10-22 00:00:00','2017-12-03 14:41:45','2017-12-03 23:41:46',NULL),
	(348067328,'etude_9','Etude 9','Hisanori Ito','','2017-10-24 00:00:00','2017-11-25 07:28:09','2017-11-25 16:28:09',NULL),
	(349016864,'blog_9','BLOG 9','抑止力 (track by Etude 9 Hisanori Ito)','何でもない様に見える、決定的なズレ、違和感の様なものを人と話してると感じる時が多々ある','2017-10-29 00:00:00','2017-12-03 14:46:12','2017-12-03 23:46:12',NULL),
	(349016865,'blog_9','BLOG 9','抑止力 (track by Etude 9 Hisanori Ito)','ブログ9こめ','2017-10-29 00:00:00','2017-11-29 01:46:35','2017-11-29 10:46:35','2017-11-29 10:46:35'),
	(349098722,'etude_10','Etude 10 ','Hisanori Ito','','2017-11-29 00:00:00','2017-12-03 12:37:10','2017-12-03 21:37:10',NULL),
	(364315028,'12_rain','12月の雨 ← New!','Ratel','久しぶりにスタジオに行って、その場で作った\r\n終わった後のオリジンダイニングは美味かったな\r\n伊藤大先生はその日2回目のオリジンだったらしい、、、\r\nこの集合写真に居る、我らがグランドコリンを見つけられた人は、三鷹のインドカレー屋で友達になろう?  (抑止力)\r\n\r\n補足すると、「集合写真」とはサウンドクラウドの Ratel アカウントの画像のことですな。 (hidelberq)','2017-12-03 12:36:00','2017-12-03 15:50:58','2017-12-04 00:50:59',NULL);

/*!40000 ALTER TABLE `soundcloud_track` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
