ALTER TABLE `trips` ADD COLUMN `user_id` BIGINT NOT NULL AFTER `name`;
ALTER TABLE `trips` ADD CONSTRAINT `fk_trip_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);