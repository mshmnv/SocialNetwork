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
 curl -X POST "http://localhost:8080/login" -d '{"id": "5", "password": "best password"}'
```

Получение анкеты пользователя
```bash
curl -X GET "http://localhost:8080/user/get/5" 
```