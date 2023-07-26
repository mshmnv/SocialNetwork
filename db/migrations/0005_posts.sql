-- +goose Up

CREATE TABLE posts (
    id bigserial PRIMARY KEY,
    text text NOT NULL,
    author_id integer REFERENCES users(id),
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    is_deleted bool DEFAULT false NOT NULL
);

-- +goose Down

DROP TABLE posts;