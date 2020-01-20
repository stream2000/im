create database account;
use account;

SET NAMES 'utf8';

#DROP TABLE IF EXISTS `account`;
CREATE TABLE IF NOT EXISTS `account` (
  `email` varchar(64) NOT NULL,
  `id` varchar(64) NOT NULL,
  `password` varchar(64) DEFAULT NULL,
  `nickname` varchar(30) NOT NULL,
  `description` varchar(64) DEFAULT NULL,
  `nickname_mtime` timestamp NULL DEFAULT NULL,
  `ctime` timestamp NULL DEFAULT NULL,
  `mtime` timestamp NULL DEFAULT NULL,
  UNIQUE KEY `account_uuid_uindex` (`id`),
  UNIQUE KEY `account_email_uindex` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;