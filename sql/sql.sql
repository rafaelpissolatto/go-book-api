-- Active: 1662320489072@@127.0.0.1@3306@devbook
CREATE DATABASE IF NOT EXISTS `devbook` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `devbook`;
DROP TABLE IF EXISTS `devbook`.`users`;
CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    nickname VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(60) NOT NULL,
    createdAt timestamp default current_timestamp(),
    PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE followers(
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;