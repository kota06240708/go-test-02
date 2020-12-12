

-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `users` DROP INDEX `icon_2`;
ALTER TABLE `users` DROP INDEX `icon_3`;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `users` ADD CONSTRAINT UNIQUE(`icon`);
