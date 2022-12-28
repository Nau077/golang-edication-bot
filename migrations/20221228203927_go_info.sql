-- +goose Up
CREATE TABLE go_info (
    id                  bigserial primary key,
    text                text,
    created_at          timestamp not null default now(),
    updated_at          timestamp
);

-- +goose Down
DROP TABLE go_info;
