# ZenTotem DSP
## Тестовое задание для iConText Group
### Консольное приложение, которое отвечает на пост-запросы

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
![Снимок экрана 2023-05-10 034041](https://github.com/VitalyMyalkin/zen_cli/assets/102473387/1bcf489f-284b-4ce1-bdc9-4dca6dbcfd9d)
2. POST `/sign/hmacsha512`
![Снимок экрана 2023-05-10 034410](https://github.com/VitalyMyalkin/zen_cli/assets/102473387/c43a6be9-378a-4cf6-9cd4-284535e60cd9)
3. POST `/postgres/users`
![Снимок экрана 2023-05-10 034021](https://github.com/VitalyMyalkin/zen_cli/assets/102473387/89233c0d-2a53-4aec-b319-e561c7206c8b)


------------
# Автор проекта:
## [Мялькин Виталий](https://github.com/VitalyMyalkin)

