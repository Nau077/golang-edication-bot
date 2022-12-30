-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE go_info ADD COLUMN text_type text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
