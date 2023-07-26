Дисклеймер: не повторяйте это на своих данных

1. Запустим Postgres:
docker run --name pg-otus -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pwd -e POSTGRES_DB=otusdb -d postgres


2. Откроем 4 подключения к БД:
docker exec -it pg-otus psql -U user -d otusdb

3. Нальем заготовку для данных:

create table test(version text primary key not null);
insert into test values('v1');

4. Проверим уровень изоляции:

select current_setting('transaction_isolation');

5. Откроем транзакцию во всех клиентах:

begin;

6. Первый клиент:
update test set version = 'v2';

commit;

7. Второй клиент:
select * from test;
update test set version = 'v3';

8. Третий клиент:

select * from test;
delete from test;


9. Второй клиент:

commit;

Разблокировался 3 клиент

10. Четвертый клиент:

select * from test;

11. Третий клиент:

commit;

12. Четвертый клиент:

select * from test;
