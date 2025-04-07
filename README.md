# Build and Run

1. Поднимаем БД через докер (опционально):
- docker compose up -d

2. Делаем миграцию:
- goose -dir ./internal/infrastructure/database/migrations postgres "host="host" user="user" database="database" password="password" sslmode=disable" up

3. Запускаем приложение:
- go run cmd/server/main.go

# Env example:
- DB_HOST=localhost
- DB_PORT=5432
- POSTGRES_USER=user
- POSTGRES_PASSWORD=password
- POSTGRES_DB=db_name
- SSL_MODE=disable
- SERVER_PORT=localhost:8080

# swagger 
- localhost:8080/swagger/index.html
