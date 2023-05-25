-- +goose Up

CREATE TABLE posts (
    id bigserial primary key,
    text text not null,
    user_id integer not null, -- todo: ?foreign key
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    is_deleted bool default false not null
);

-- +goose Down

DROP TABLE posts;