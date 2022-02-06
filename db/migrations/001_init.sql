-- +goose Up

CREATE TABLE users (
    id uuid PRIMARY KEY,
    name character varying not null default '',
    created_at timestamp without time zone not null default now()
);

-- +goose Down
-- DROP SCHEMA public CASCADE;