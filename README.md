# HTTP API для асинхронных задач

## Описание
Простой сервис на Go для запуска долгих задач и получения их результатов через HTTP API.

## Запуск проекта

```bash
docker-compose up --build
```

Сервис будет доступен на:
```
http://localhost:8080
```

## API Эндпоинты

### POST /tasks
Создание новой задачи.

**Пример запроса:**
```bash
curl -X POST http://localhost:8080/tasks
```

**Ответ:**
```json
{
  "task_id": "your-task-id"
}
```

### GET /tasks/{id}
Получение статуса и результата задачи.

**Пример запроса:**
```bash
curl http://localhost:8080/tasks/{task_id}
```

**Ответы:**
- Статус задачи: `pending`, `in_progress`, `done`
- Результат задачи при готовности

## Структура проекта

```
cmd/            // Точка входа в приложение
internal/
  app/          // Инициализация сервера и роутинг
  task/         // Логика задач
  worker/       // Фоновые воркеры
pkg/
  utils/        // Утилиты
Dockerfile      // Сборка образа
docker-compose.yml // Оркестрация
go.mod
README.md
```

## Требования

- Go 1.22+
- Docker
