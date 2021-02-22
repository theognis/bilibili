CREATE DATABASE IF NOT EXISTS bilibili;
USE bilibili;

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo`
(
    `uid`                INT AUTO_INCREMENT PRIMARY KEY,
    `username`           VARCHAR(15)  NOT NULL,
    `gender`             VARCHAR(1)   NOT NULL DEFAULT 'N',
    `phone`              VARCHAR(11)  NOT NULL,
    `salt`               VARCHAR(10)  NOT NULL,
    `password`           VARCHAR(32)  NOT NULL,
    `email`              VARCHAR(20)  NOT NULL DEFAULT '',
    `statement`          VARCHAR(90)  NOT NULL DEFAULT '这个人很懒，什么都没有写',
    `avatar`             VARCHAR(120) NOT NULL DEFAULT 'https://redrock.oss-cn-chengdu.aliyuncs.com/akari.jpg',
    `reg_date`           DATE         NOT NULL,
    `birthday`           DATE         NOT NULL DEFAULT '9999-12-12',
    `last_check_in_date` DATE         NOT NULL DEFAULT '1926-08-17',
--    `last_coin_date`     DATE         NOT NULL DEFAULT '1926-08-17',
--    `daily_coin`         INT          NOT NULL DEFAULT 0,
--    `last_view_date`     DATE         NOT NULL DEFAULT '1926-08-17',
--    `daily_view`         INT          NOT NULL DEFAULT 0,
    `exp`                INT          NOT NULL DEFAULT 0,
    `coins`              INT          NOT NULL DEFAULT 0,
    `b_coins`            INT          NOT NULL DEFAULT 0,
    UNIQUE (`username`),
    UNIQUE (`phone`)
) charset="utf8mb4";

alter table userinfo add daily_coin int not null default 0;
alter table userinfo add last_coin_date date not null default '1926-08-17';
alter table userinfo add daily_view int not null default 0;
alter table userinfo add last_view_date date not null default '1926-08-17';

DROP TABLE IF EXISTS `video_label`;

CREATE TABLE `video_label`
(
    `id`          INT AUTO_INCREMENT PRIMARY KEY,
    `av`          INT         NOT NULL,
    `video_label` VARCHAR(10) NOT NULL
) charset="utf8mb4";

DROP TABLE IF EXISTS `video_like`;

CREATE TABLE `video_like`
(
    `id`  INT AUTO_INCREMENT PRIMARY KEY,
    `av`  INT NOT NULL,
    `uid` INT NOT NULL
) charset="utf8mb4";

DROP TABLE IF EXISTS `video_coin`;

CREATE TABLE `video_coin`
(
    `id`  INT AUTO_INCREMENT PRIMARY KEY,
    `av`  INT NOT NULL,
    `uid` INT NOT NULL
) charset="utf8mb4";

DROP TABLE IF EXISTS `video_save`;

CREATE TABLE `video_save`
(
    `id`  INT AUTO_INCREMENT PRIMARY KEY,
    `av`  INT NOT NULL,
    `uid` INT NOT NULL
) charset="utf8mb4";

DROP TABLE IF EXISTS `video_info`;

CREATE TABLE `video_info`
(
    `av`          INT AUTO_INCREMENT PRIMARY KEY,
    `title`       VARCHAR(80)  NOT NULL,
    `channel`     VARCHAR(4)   NOT NULL,
    `description` VARCHAR(250) NOT NULL,
    `video_url`   VARCHAR(120) NOT NULL,
    `cover_url`   VARCHAR(120) NOT NULL,
    `author_uid`  INT          NOT NULL,
    `time`        TIMESTAMP    NOT NULL,
    `views`       INT          NOT NULl DEFAULT 0,
    `likes`       INT          NOT NULL DEFAULT 0,
    `coins`       INT          NOT NULL DEFAULT 0,
    `saves`       INT          NOT NULL DEFAULT 0,
    `shares`      INT          NOT NULL DEFAULT 0
) charset="utf8mb4";

DROP TABLE IF EXISTS `video_danmaku`;

CREATE TABLE `video_danmaku`
(
    `did`      INT AUTO_INCREMENT PRIMARY KEY,
    `av`       INT          NOT NULL,
    `uid`      INT          NOT NULL,
    `value`    VARCHAR(120) NOT NULL,
    `color`    VARCHAR(6)   NOT NULL,
    `type`     VARCHAR(10)  NOT NULL,
    `time`     TIMESTAMP    NOT NULL,
    `location` INT          NOT NULL
) charset="utf8mb4";
