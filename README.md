# SocialNetwork

Запуск
```bash
  make up
```

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
curl -X GET "http://localhost:8080/user/get/5" 
```

Поиск анкет пользователей
```bash
curl -X GET "http://localhost:8080/user/search?first_name=A&second_name=A" 
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