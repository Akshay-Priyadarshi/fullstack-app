package server

import (
	"fmt"
	"log/slog"

	"github.com/Akshay-Priyadarshi/fullstack-app/db/connections"
	"github.com/Akshay-Priyadarshi/fullstack-app/internal/app/server/initialisations"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Port      string `json:"port"`
	DBString  string `json:"dbString"`
	ApiPath   string `json:"apiPath"`
	WebPath   string `json:"webPath"`
	JwtSecret string `json:"jwtSecret"`
	Env       string `json:"env"`
}

type Server struct {
	App        *fiber.App
	Config     *Config
	DB         *sqlx.DB
	Logger     *slog.Logger
	Translator *ut.Translator
	Validate   *validator.Validate
}

func New(app *fiber.App, config *Config, logger *slog.Logger) *Server {
	slog.Debug("Application configuration", "config", config)
	return &Server{
		App:    app,
		Config: config,
		DB:     connections.InitializeDB(config.DBString),
		Logger: logger,
	}
}

func (s *Server) Start() error {
	slog.Info("Server is starting at: http://localhost:" + s.Config.Port)
	if err := s.App.Listen(fmt.Sprintf(":%s", s.Config.Port)); err != nil {
		return err
	}
	return nil
}

func (s *Server) Configure() {
	// Register recover middleware
	s.App.Use(recover.New())

	// Initialisations
	s.Translator = initialisations.InitI18n()
	s.Validate = initialisations.InitValidation(s.Translator)

	// Enable CORS
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Enable security headers
	s.App.Use(helmet.New())

	// Enable swagger api documentation
	s.App.Get("/swagger/*", swagger.HandlerDefault)
}

var AppServer *Server
