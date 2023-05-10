# ZenTotem DSP
## Тестовое задание для iConText Group
### Консольное приложение, которое отвечает на пост-запросы

------------
# Автор проекта:
### [Мялькин Виталий](https://github.com/VitalyMyalkin)

------------

### Запуск проекта

1. Клонируйте репозиторий

2. В терминале выполните команду для запуска сервера 

```
go run main.go
```
3. В терминале введите хост и порт подключения к Redis, например:
```
Введите хост подключения к Redis
localhost
Введите порт подключения к Redis
6379
```
4. Запись в Postgres происходит в базу данных zen, хост localhost, порт 5433
------------

### Список и примеры использования эндпоинтов (согласно тестовому заданию):

1. POST `/redis/incr`

2. POST `/sign/hmacsha512`

3. POST `/postgres/users`
