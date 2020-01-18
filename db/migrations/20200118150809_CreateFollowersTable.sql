
-- +goose Up
CREATE TABLE IF NOT EXISTS `followers` (
  `follower_id` INT UNSIGNED NOT NULL,
  `followee_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  INDEX `fk_followers_users_idx` (`follower_id` ASC) VISIBLE,
  INDEX `fk_followers_users1_idx` (`followee_id` ASC) VISIBLE,
  PRIMARY KEY (`follower_id`, `followee_id`),
  CONSTRAINT `fk_followers_users`
    FOREIGN KEY (`follower_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_followers_users1`
    FOREIGN KEY (`followee_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
);


-- +goose Down
DROP TABLE `followers`;

