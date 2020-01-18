
-- +goose Up
CREATE TABLE IF NOT EXISTS `likes` (
  `tweets_id` INT UNSIGNED NOT NULL,
  `users_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`tweets_id`, `users_id`),
  INDEX `fk_likes_users1_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_likes_tweets1`
    FOREIGN KEY (`tweets_id`)
    REFERENCES `tweets` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_likes_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);


-- +goose Down
DROP TABLE `likes`;
