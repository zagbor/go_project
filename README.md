# Go Microservice Homework

## Описание
Высоконагруженный микросервис на Go для управления пользователями.

## Функциональность
- CRUD операции пользователей (`/api/users`).
- Rate Limiting (1000 req/s).
- Prometheus Metrics (`/metrics`).
- Асинхронное логирование.

## Запуск

### Локально
```bash
go mod tidy
go run main.go
```

### Через Docker Compose
```bash
```

## Тестирование
Для нагрузочного тестирования используйте `wrk`:
```bash
wrk -t12 -c500 -d60s http://localhost:8080/api/users
```

## Структура
- `main.go`: Точка входа.
- `handlers`: HTTP обработчики.
- `services`: Бизнес-логика.
- `utils`: Логгер, Rate Limiter.
- `metrics`: Prometheus.
