-- +goose Up
-- SQL in this section is executed when the migration is applied.

create role replicator with login replication password 'replicator_password';