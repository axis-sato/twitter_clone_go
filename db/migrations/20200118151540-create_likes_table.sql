
-- +migrate Up
CREATE TABLE IF NOT EXISTS `likes` (
  `tweet_id` INT UNSIGNED NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`tweet_id`, `user_id`),
  INDEX `fk_likes_users1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_likes_tweets1`
    FOREIGN KEY (`tweet_id`)
    REFERENCES `tweets` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_likes_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);


-- +migrate Down
DROP TABLE `likes`;
