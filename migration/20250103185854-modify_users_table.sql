
-- +migrate Up

ALTER TABLE `users`
MODIFY `role` ENUM('librarian', 'visitor', 'owner') NOT NULL;

ALTER TABLE `users`
ADD COLUMN `username` VARCHAR(255) NOT NULL AFTER `email`;


-- +migrate Down

ALTER TABLE `users`
MODIFY `role` ENUM('librarian', 'visitor') NOT NULL;

ALTER TABLE `users`
DROP COLUMN `username`;
