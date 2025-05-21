# Мини-сервис “Цитатник”

## 📌 Описание проекта

Мини-сервис на **Go** для хранения и управления цитатами.  
Использует **REST API** и минимальные зависимости (`gorilla/mux`).  
Позволяет **добавлять, получать, фильтровать и удалять цитаты**.

## 🚀 Функциональность

- **Добавить цитату:** `POST /quotes`
- **Получить все цитаты:** `GET /quotes`
- **Получить случайную цитату:** `GET /quotes/random`
- **Фильтровать по автору:** `GET /quotes?author=Confucius`
- **Удалить цитату:** `DELETE /quotes/{id}`

## 🚀 Запуск приложения

`go run main.go`

## 🚀 Запуск тестов

`go test ./... -v`

## CURL-команды

- Создать цитату:
  `curl -X POST http://localhost:8080/quotes \
 -H "Content-Type: application/json"\
  -d  '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
`

- Создать цитату:`curl http://localhost:8080/quotes `

- Создать цитату:`curl http://localhost:8080/quotes/random`

- Создать цитату:`curl http://localhost:8080/quotes?author=Confucius`

- Создать цитату:`curl -X DELETE http://localhost:8080/quotes/1`
