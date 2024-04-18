-- `imersao-go`.files definition

CREATE TABLE `files` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `folder_id` int(11) DEFAULT NULL,
  `owner_id` int(11) NOT NULL,
  `name` varchar(200) NOT NULL,
  `type` varchar(50) NOT NULL,
  `path` varchar(250) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modified_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `files_folders_FK` (`folder_id`),
  KEY `files_users_FK` (`owner_id`),
  CONSTRAINT `files_folders_FK` FOREIGN KEY (`folder_id`) REFERENCES `folders` (`id`),
  CONSTRAINT `files_users_FK` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;