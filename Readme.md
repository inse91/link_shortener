# Укорачиватель ссылок

### grpc & http сервис преобразования ссылок

### Запуск

`go run cmd/main.go` - local
`shortener:latest` - docker image

#### Переменные окружения

|key|default|
|---|---|
|GRPC_PORT| 10000|
|HTTP_PORT| 10010|
|DB_CONNECTION||

Если `DB_CONNECTION` не указывается то используется inMemory хранилище

#### API

`api/swagger.json` - openAPI 2.0
`api/shorter.proto` - proto