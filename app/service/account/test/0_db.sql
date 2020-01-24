create database account;
use account;

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account` (
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `email` varchar(64) NOT NULL,
                           `password` varchar(50) DEFAULT NULL,
                           `nickname` varchar(64) NOT NULL,
                           `nickname_mtime` timestamp NULL DEFAULT NULL,
                           `ctime` timestamp NULL DEFAULT NULL,
                           `mtime` timestamp NULL DEFAULT NULL,
                           `sign` varchar(30) DEFAULT '这个人很懒，什么都没有留下',
                           `profile_pic_url` varchar(100) DEFAULT 'default',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `account_email_uindex` (`email`)
);
