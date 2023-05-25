# SocialNetwork

Запуск
```bash
  make up
  make migration-up
```

Регистрация пользователя
```bash
 curl -X POST "http://localhost:8080/user/register" -d '{"first_name": "Sam", "second_name": "Sim", "age": 60, "birthdate": "1970-08-15", "biography": "love cats and dogs", "city": "Rome", "password": "best password"}'
```

Авторизация пользователя
```bash
 curl -X POST "http://localhost:8080/login" -d '{"user_id": "1009900", "password": "best password"}'
```

Получение анкеты пользователя
```bash
curl -X GET "http://localhost:8080/user/get/1009900"
```

Поиск анкет пользователей
```bash
curl -X GET "http://localhost:8080/user/search?first_name=A&second_name=A" 
```

Отправить заявку в друзья
```bash
curl -X PUT "http://localhost:8080/friend/set/1009902"  -d "{}" --cookie  "session-token=b861fd01-9336-4f5d-a0cd-bb6aeca5df29"
```

Удалить друга
```bash
curl -X PUT "http://localhost:8080/friend/delete/1009902"  -d "{}" --cookie  "session-token=9ba5bdd1-8f6b-42ec-9636-caf0f05ae14d"
```


Заполнение базы данными 
```bash
curl -X POST "http://localhost:8080/add-users" 
```

---

### *Highload Architect Course:*
1. MVP
2. Indexes and Load Testing
3. Replication
4. Cache
5. 