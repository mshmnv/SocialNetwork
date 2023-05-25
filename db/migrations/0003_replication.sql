-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE ROLE replicator WITH LOGIN REPLICATION PASSWORD 'replicator_password';
