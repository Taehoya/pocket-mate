CREATE TABLE IF NOT EXISTS `trip` (
    `trip_idx` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(30),
    `budget` DECIMAL(13,4),
    `country_id` BIGINT,
    `description` VARCHAR(255),
    `start_date_time` DATETIME NOT NULL,
    `end_date_time` DATETIME NOT NULL,
    PRIMARY KEY (trip_idx)
)