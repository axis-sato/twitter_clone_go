
-- +goose Up
CREATE TABLE IF NOT EXISTS `tweets` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `users_id` INT UNSIGNED NOT NULL,
  `tweet` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `update_at` TIMESTAMP NOT NULL,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_tweets_users1_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_tweets_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);


-- +goose Down
DROP TABLE `tweets`;
