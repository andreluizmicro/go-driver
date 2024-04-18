-- `imersao-go`.folders definition

CREATE TABLE `folders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `name` varchar(60) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `folders_folders_FK` (`parent_id`),
  CONSTRAINT `folders_folders_FK` FOREIGN KEY (`parent_id`) REFERENCES `folders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;