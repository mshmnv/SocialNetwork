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