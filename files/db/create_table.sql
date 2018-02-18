CREATE DATABASE restohub;
USE restohub;

-- Adminer 4.3.1 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `restaurant_id` int(11) NOT NULL,
  `book_datetime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `total_people` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_BOOK_USER` (`user_id`),
  KEY `FK_BOOK_RESTO` (`restaurant_id`),
  CONSTRAINT `FK_BOOK_RESTO` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants` (`id`),
  CONSTRAINT `FK_BOOK_USER` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `countries`;
CREATE TABLE `countries` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `countries` (`id`, `name`) VALUES
(1,	'Indonesia');

DROP TABLE IF EXISTS `cuisines`;
CREATE TABLE `cuisines` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `cuisines` (`id`, `name`) VALUES
(1,	'Chinese');

DROP TABLE IF EXISTS `restaurants`;
CREATE TABLE `restaurants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `slug` varchar(50) NOT NULL,
  `cuisine_id` int(11) NOT NULL,
  `country_id` int(11) NOT NULL,
  `lat` float(10,6) NOT NULL,
  `lng` float(10,6) NOT NULL,
  `address` text,
  `rating` float(2,1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_RESTO_CUISINE` (`cuisine_id`),
  KEY `FK_RESTO_COUNTRY` (`country_id`),
  CONSTRAINT `FK_RESTO_COUNTRY` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`),
  CONSTRAINT `FK_RESTO_CUISINE` FOREIGN KEY (`cuisine_id`) REFERENCES `cuisines` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `restaurants` (`id`, `name`, `slug`, `cuisine_id`, `country_id`, `lat`, `lng`, `address`, `rating`) VALUES
(1,	'Holycow',	'holycow',	1,	1,	1.000000,	1.000000,	'Jl. Guru Mughni No.12',	5.0);

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `email` varchar(320) NOT NULL,
  `password` char(60) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(1,	'antony',	'antonybudianto@gmail.com',	'1234');

-- 2018-02-18 14:36:30