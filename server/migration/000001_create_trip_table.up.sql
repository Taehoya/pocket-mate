CREATE TABLE IF NOT EXISTS `trips` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(30),
    `budget` DECIMAL(13,4),
    `country_id` BIGINT,
    `description` VARCHAR(255),
    `start_date_time` DATETIME NOT NULL,
    `end_date_time` DATETIME NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL, 
    PRIMARY KEY (`id`)
)