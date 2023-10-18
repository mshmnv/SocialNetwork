-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE messaging (
    dialog_id bigserial NOT NULL,
    sender_id integer NOT NULL,
    receiver_id integer NOT NULL,
    sent_at  timestamp DEFAULT now() NOT NULL,
    text text NOT NULL
);

CREATE INDEX messaging_dialog_id ON messaging (dialog_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE messaging;