CREATE DATABASE IF NOT EXISTS bilibili;
USE bilibili;

DROP TABLE IF EXISTS `userinfo`;

CREATE TABLE `userinfo` (
    `uid` INT AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(15) NOT NULL,
    `phone` VARCHAR(11) NOT NULL, 
    `salt` VARCHAR(10) NOT NULL,
    `password` VARCHAR(32) NOT NULL,
    `reg_date` DATE NOT NULL,
    `email` VARCHAR(20) NOT NULL DEFAULT "",
    `statement` VARCHAR(90) DEFAULT '这个人很懒，什么都没有写',
    UNIQUE (`username`),
    UNIQUE (`phone`)
);
