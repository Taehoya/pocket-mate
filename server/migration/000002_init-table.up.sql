-- create countries table
CREATE TABLE IF NOT EXISTS `countries` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(10) DEFAULT NULL,
  `name` VARCHAR(50) DEFAULT NULL,
  `currency` VARCHAR(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
);

-- create categories table
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

-- create users table
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `nickname` VARCHAR(30) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL, 
    PRIMARY KEY (`id`),
    UNIQUE KEY (`nickname`)
);

-- create transactions table
CREATE TABLE IF NOT EXISTS `transactions` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `trip_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  `name` VARCHAR(30) DEFAULT NULL,
  `amount` DECIMAL(13,4) DEFAULT NULL,
  `date` DATETIME NOT NULL,
  `category_id` BIGINT DEFAULT NULL,
  `description` VARCHAR(255) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `FK_TransactionTrip` (`trip_id`),
  KEY `FK_TransactionCategory` (`category_id`),
  CONSTRAINT `FK_TransactionCategory` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `FK_TransactionTrip` FOREIGN KEY (`trip_id`) REFERENCES `trips` (`id`) ON DELETE CASCADE
);


-- create payments table
CREATE TABLE IF NOT EXISTS `payments` (
  `transaction_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  `status` TINYINT(1) DEFAULT NULL,
  PRIMARY KEY (`transaction_id`,`user_id`),
  KEY `FK_PayerUser` (`user_id`),
  CONSTRAINT `FK_PayerTransaction` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FK_PayerUser` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);


-- create user_trips table
CREATE TABLE IF NOT EXISTS `user_trips` (
  `leader` TINYINT(1) DEFAULT NULL,
  `trip_id` BIGINT NOT NULL,
  `user_id` BIGINT NOT NULL,
  PRIMARY KEY (`trip_id`,`user_id`),
  KEY `FK_UserTrip_User` (`user_id`),
  CONSTRAINT `FK_UserTrip_Trip` FOREIGN KEY (`trip_id`) REFERENCES `trips` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FK_UserTrip_User` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
);
