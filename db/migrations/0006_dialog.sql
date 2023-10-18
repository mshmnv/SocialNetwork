-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE dialog (
    id bigserial PRIMARY KEY NOT NULL,
    sender_id integer NOT NULL REFERENCES users (id),
    receiver_id integer NOT NULL REFERENCES users (id)
);

CREATE INDEX dialog_index ON dialog (sender_id, receiver_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE dialog;