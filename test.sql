DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `category` (`id`, `name`, `created_at`, `updated_at`)
VALUES
	(1,'计算机','2018-01-20 11:31:43','2018-01-20 11:31:43'),
	(2,'经济学','2018-01-20 11:31:47','2018-01-20 11:31:47'),
	(3,'心理学','2018-01-20 11:31:51','2018-01-20 11:31:51'),
	(4,'经济学','2018-01-20 11:32:02','2018-02-06 14:54:22');


DROP TABLE IF EXISTS `publisher`;

CREATE TABLE `publisher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `addr` varchar(100) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `publisher` (`id`, `name`, `addr`, `created_at`, `updated_at`)
VALUES
	(1,'清华大学出版社','北京市海淀区','2018-01-20 11:24:39','2018-01-20 11:25:54'),
	(2,'人民出版社','北京市海淀区','2018-01-20 11:26:30','2018-01-20 11:28:44'),
	(3,'中信出版社','北京市海淀区','2018-01-20 11:27:54','2018-01-20 11:29:11'),
	(4,'高等教育出版社','北京市海淀区','2018-01-20 11:28:54','2018-01-20 11:29:12'),
	(5,'工信出版社','北京市海淀区','2018-01-20 11:29:03','2018-02-06 14:54:22'),
	(6,'工信出版社','重庆市','2018-01-20 11:29:46','2018-02-06 11:52:05');
