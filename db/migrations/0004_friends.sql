-- +goose Up

CREATE TYPE friendship_status AS ENUM ('pending', 'approved');

CREATE TABLE friends (
  user_id integer NOT NULL REFERENCES users (id),
  friend_id integer NOT NULL REFERENCES users (id),
  status friendship_status NOT NULL,
  UNIQUE (user_id, friend_id)
);

-- +goose Down

DROP TABLE friends;