-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE INDEX first_second_name_index ON users USING btree(UPPER(first_name) text_pattern_ops, UPPER(second_name) text_pattern_ops);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP INDEX first_second_name_index;
