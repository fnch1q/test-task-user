// @title API
// @version 1.0
// @description API for test task
// @BasePath /
package main

import (
	"test-task-user/internal/config"
	"test-task-user/internal/infrastructure"
	"time"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	_ = cfg

	app := infrastructure.NewApp(cfg).
		ContextTimeout(5 * time.Second).
		SetupLogger().
		SetupDB().
		SetupEnrichment().
		SetupServer()

	app.Start()
}

/*
2. Создать entity User (национальность text, age смотрит на probability (ставится в .env)), repo User
3. Сделать миграцию (USER
	name VARCHAR(50) NOT NULL,
	surname VARCHAR(50) NOT NULL,
	patronymic VARCHAR(50) NULL,
	age INT NOT NULL,
	gender enum NOT NULL,
	nationality text NOT NULL,
)
4. usecase: GET, POST, PUT, DELETE
*/
