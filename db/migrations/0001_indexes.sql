-- +goose Up
-- SQL in this section is executed when the migration is applied.

create index first_second_name_index on users using btree(UPPER(first_name) text_pattern_ops, UPPER(second_name) text_pattern_ops);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop index first_second_name_index;
