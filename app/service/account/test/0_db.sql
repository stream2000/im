create database account;
use account;

SET NAMES 'utf8';

#DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
  `email` varchar(256) NOT NULL,
  `id` varchar(256) NOT NULL,
  `password` varchar(256) DEFAULT NULL,
  `nickname` varchar(100) NOT NULL,
  `nickname_mtime` timestamp NULL DEFAULT NULL,
  `ctime` timestamp NULL DEFAULT NULL,
  `mtime` timestamp NULL DEFAULT NULL,
  `profile_pic_url` varchar(30) DEFAULT '这个人很懒，什么都没有留下',
  `sign` varchar(30) DEFAULT NULL,
  UNIQUE KEY `account_uuid_uindex` (`id`),
  UNIQUE KEY `account_email_uindex` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;