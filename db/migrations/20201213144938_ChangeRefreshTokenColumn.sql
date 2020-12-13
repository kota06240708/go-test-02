
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `refresh_tokens` MODIFY `expire` DATETIME NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `refresh_tokens` MODIFY `expire` VARCHAR(255) NOT NULL;
