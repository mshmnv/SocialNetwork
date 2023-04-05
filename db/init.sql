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

INSERT INTO users (first_name, second_name, age, birthdate, biography, city) VALUES ('Ivan', 'Ivanovich',  30, '1995-12-31', 'volleyball', 'Moscow');
INSERT INTO users (first_name, second_name, age, birthdate, biography, city) VALUES ('Sofia', 'An',  18, '2005-08-15', 'birds', 'Nizhniy');
INSERT INTO users (first_name, second_name, age, birthdate, biography, city) VALUES ('Max', 'Sinizyn',  27, '1980-01-22', 'football', 'Moscow');
INSERT INTO users (first_name, second_name, age, birthdate, biography, city) VALUES ('Petr', 'Petrovich', 33, '1999-10-26', 'book, newspapers, magazines', 'Saint Petersburg');
