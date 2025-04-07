package infrastructure

import (
	"test-task-user/internal/config"
	"time"

	db "test-task-user/internal/infrastructure/database"
	"test-task-user/internal/infrastructure/enrichment"
	server "test-task-user/internal/infrastructure/router"

	"github.com/sirupsen/logrus"
)

type app struct {
	dbManager  db.DBManager
	cfg        config.Config
	logger     *logrus.Logger
	server     server.Server
	enrichment enrichment.Client
	ctxTimeout time.Duration
}

func NewApp(config config.Config) *app {
	return &app{
		cfg: config,
	}
}

func (a *app) ContextTimeout(t time.Duration) *app {
	a.ctxTimeout = t
	return a
}

func (a *app) SetupLogger() *app {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(logrus.InfoLevel)
	a.logger = logger
	a.logger.Infof("Success log configured")
	return a
}

func (a *app) SetupDB() *app {
	gormDB, err := db.NewPostgresConnection(a.cfg)
	if err != nil {
		a.logger.Fatalln("Failed to setup database:", err)
	}
	a.dbManager = *db.NewDBManager(gormDB)
	a.logger.Infof("Success setup database")
	return a
}

func (a *app) SetupEnrichment() *app {
	a.enrichment = enrichment.NewClient()
	return a
}

func (a *app) SetupServer() *app {
	a.server = server.NewServer(a.cfg, a.logger, a.dbManager, a.enrichment, a.ctxTimeout)
	a.logger.Infof("Success setup server")
	return a
}

func (a *app) Start() {
	a.server.Listen()
}
