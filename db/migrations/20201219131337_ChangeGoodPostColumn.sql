
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `goods` CHANGE COLUMN `isGood` `is_good` BOOLEAN;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `goods` CHANGE COLUMN `is_good` `isGood` BOOLEAN;
