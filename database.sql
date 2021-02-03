CREATE
DATABASE IF NOT EXISTS bilibili;
USE
bilibili;

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo`
(
    `uid`                INT AUTO_INCREMENT PRIMARY KEY,
    `username`           VARCHAR(15)  NOT NULL,
    `gender`             VARCHAR(1)   NOT NULL DEFAULT 'N', --N为未知，F女性，M男性，O其它。
    `phone`              VARCHAR(11)  NOT NULL,
    `salt`               VARCHAR(10)  NOT NULL,
    `password`           VARCHAR(32)  NOT NULL,
    `email`              VARCHAR(20)  NOT NULL DEFAULT '',
    `statement`          VARCHAR(90)  NOT NULL DEFAULT '这个人很懒，什么都没有写',
    `avatar`             VARCHAR(120) NOT NULL DEFAULT 'https://redrock.oss-cn-chengdu.aliyuncs.com/akari.jpg',
    `reg_date`           DATE         NOT NULL,
    `birthday`           DATE         NOT NULL DEFAULT '9999-12-12', --该默认值应该识别为未知
    `last_check_in_date` DATE         NOT NULL DEFAULT '1926-08-17',
    `exp`                INT          NOT NULL DEFAULT 0,
    `coins`              INT          NOT NULL DEFAULT 0,
    `b_coins`            INT          NOT NULL DEFAULT 0,
    UNIQUE (`username`),
    UNIQUE (`phone`)
) charset="utf8mb4";
