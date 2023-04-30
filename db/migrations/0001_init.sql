-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE users (
    id bigserial primary key,
    first_name text not null,
    second_name text not null,
    age integer default 0 not null,
    birthdate text not null,
    biography text not null,
    city text not null,
    password text default '' not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users;
