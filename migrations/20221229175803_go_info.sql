-- +goose Up
-- +goose StatementBegin
ALTER TABLE go_info ADD COLUMN title text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
