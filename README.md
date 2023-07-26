# SocialNetwork

Запуск
```bash
  make up
  make migration-up
```

## Postman Коллекция
/testing/postman_collection.json

## USERS

Регистрация пользователя
```bash
 curl -X POST "http://localhost:8080/user/register" -d '{"first_name": "Sam", "second_name": "Sim", "age": 60, "birthdate": "1970-08-15", "biography": "love cats and dogs", "city": "Rome", "password": "best password"}'
```

Авторизация пользователя
```bash
 curl -X POST "http://localhost:8080/login" -d '{"id": "1009900", "password": "best password"}'
```

Получение анкеты пользователя
```bash
curl -X GET "http://localhost:8080/user/get/1009900"
```

Поиск анкет пользователей
```bash
curl -X GET "http://localhost:8080/user/search?first_name=A&second_name=A" 
```

Заполнение базы данными
```bash
curl -X POST "http://localhost:8080/add-users" 
```

## FRIENDS

Отправить заявку в друзья
```bash
curl -X PUT "http://localhost:8080/friend/set/1009900"  --cookie  "session-token=fa1fb628-d13d-433f-bb42-dad17c7e4c07"
```

Удалить друга
```bash
curl -X PUT "http://localhost:8080/friend/delete/1009902" --cookie  "session-token=9ba5bdd1-8f6b-42ec-9636-caf0f05ae14d"
```

## POSTS
Создать пост
```bash
curl -X POST "http://localhost:8080/post/create" -d '{"text": "My super cool post"}' --cookie  "session-token=5aad114d-710d-4c17-8a77-70b81c733cc8"
```
Обновить пост
```bash
curl -X POST "http://localhost:8080/post/update" -d '{"id":8, "text": "My super cool updated post"}' --cookie  "session-token=dba9660f-76c2-4904-80e4-c6534b9bde70"

```
Удалить пост
```bash
curl -X PUT "http://localhost:8080/post/delete/1" --cookie  "session-token=03c92275-a291-4c72-a7ed-b5352cfe6de0"

```
Получить пост
```bash
curl -X GET "http://localhost:8080/post/get/1"
```

Лента постов друзей
```bash
curl -X GET "http://localhost:8080/post/feed"  --cookie  "session-token=4e6ad747-ad98-4e5e-bde4-f741d81104e5"

```


---

### *Highload Architect Course:*
1. MVP
2. Indexes and Load Testing
3. Replication
4. Cache [x]
5. Queues
6. In-memory db
7. Sharding
8. Microservices
9. Balancing and Fault tolerance
10. Distributed Transactions