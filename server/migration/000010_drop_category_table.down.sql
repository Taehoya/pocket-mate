CREATE TABLE IF NOT EXISTS `categories` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `trip_id` BIGINT NOT NULL,
  `name` VARCHAR(30) DEFAULT NULL,
  `colour` VARCHAR(30) DEFAULT NULL,
  `icon` VARCHAR(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_CategoryTrip` (`trip_id`),
  CONSTRAINT `FK_CategoryTrip` FOREIGN KEY (`trip_id`) REFERENCES `trips` (`id`) ON DELETE CASCADE
);