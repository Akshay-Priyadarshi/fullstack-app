package server

import (
	"fmt"

	"github.com/Akshay-Priyadarshi/fullstack-app/db/connections"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port      string
	DBString  string
	ApiPath   string
	WebPath   string
	JwtSecret string
}

type Server struct {
	App    *fiber.App
	DB     *sqlx.DB
	Config *Config
}

func New(app *fiber.App, config *Config) *Server {
	return &Server{
		App:    app,
		DB:     connections.InitializeDB(config.DBString),
		Config: config,
	}
}

func (s *Server) Start() error {
	s.Configure()
	println("Server is starting at: http://localhost:" + s.Config.Port)
	if err := s.App.Listen(fmt.Sprintf(":%s", s.Config.Port)); err != nil {
		return err
	}
	return nil
}

func (s *Server) Configure() {

	// Enable CORS
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Enable security headers
	s.App.Use(helmet.New())

	// Enable logging
	s.App.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Enable swagger api documentation
	s.App.Get("/swagger/*", swagger.HandlerDefault)
}

var AppServer *Server
