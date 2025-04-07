package server

import (
	"net/http"
	"test-task-user/internal/config"
	"test-task-user/internal/infrastructure/database"
	"test-task-user/internal/infrastructure/enrichment"
	"time"

	_ "test-task-user/docs" // Пакет, содержащий сгенерированный код Swagger

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/negroni"
)

type Server struct {
	cfg        config.Config
	log        *logrus.Logger
	db         database.DBManager
	enrichment enrichment.Client
	ctxTimeout time.Duration
	router     *gin.Engine
}

func NewServer(
	cfg config.Config,
	log *logrus.Logger,
	db database.DBManager,
	enrichment enrichment.Client,
	ctxTimeout time.Duration,
) Server {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	return Server{
		cfg:        cfg,
		log:        log,
		db:         db,
		enrichment: enrichment,
		ctxTimeout: ctxTimeout,
		router:     router,
	}
}

func (s *Server) Listen() {
	gin.SetMode(gin.ReleaseMode)

	s.setupRoutes()

	n := negroni.New()
	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowPrivateNetwork: true,
		AllowedMethods:      []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials:    true,
	})

	n.Use(c)
	n.UseHandler(s.router)

	server := &http.Server{
		Addr:         s.cfg.ServerPort,
		Handler:      n,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.log.Printf("Starting server on port %s", s.cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Fatalf("Server failed: %s\n", err)
	}
}

func (s *Server) setupRoutes() {
	s.router.POST("/user", s.buildCreateUser())
	s.router.DELETE("/user/:id", s.buildDeleteUser())
	s.router.PUT("/user/:id", s.buildUpdateUser())
	s.router.GET("/user/:id", s.buildGetUserByID())
	s.router.GET("/user", s.buildGetAllUsers())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}

func buildParams(c *gin.Context, params ...string) {
	q := c.Request.URL.Query()

	for _, value := range params {
		switch value {
		case "page":
			if _, exists := q["page"]; !exists {
				q.Set("page", "1")
			}
		case "limit":
			if _, exists := q["limit"]; !exists {
				q.Set("limit", "10")
			}
		default:
			q.Add(value, c.Param(value))
		}
	}
	c.Request.URL.RawQuery = q.Encode()
}

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}
