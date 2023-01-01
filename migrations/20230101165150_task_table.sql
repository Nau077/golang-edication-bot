-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id                  bigserial primary key,
    title               text,
    tasks_text          text,
    tasks_solution      text,
    created_at          timestamp not null default now(),
    updated_at          timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
