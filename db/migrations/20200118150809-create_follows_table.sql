
-- +migrate Up
CREATE TABLE IF NOT EXISTS `follows` (
  `follower_id` INT UNSIGNED NOT NULL,
  `followee_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  INDEX `fk_followers_users_idx` (`follower_id` ASC) VISIBLE,
  INDEX `fk_followees_users_idx` (`followee_id` ASC) VISIBLE,
  PRIMARY KEY (`follower_id`, `followee_id`),
  CONSTRAINT `fk_followers_users`
    FOREIGN KEY (`follower_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_followees_users`
    FOREIGN KEY (`followee_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);


-- +migrate Down
DROP TABLE `follows`;

