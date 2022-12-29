-- +goose Up
-- +goose StatementBegin
ALTER TABLE go_info 
ADD CONSTRAINT AK_go_info UNIQUE (title);   
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
