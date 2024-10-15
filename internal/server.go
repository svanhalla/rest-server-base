package internal

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/svanhalla/base-rest-server/api"
	"github.com/svanhalla/base-rest-server/templates"
)

// Server implementerar det genererade gränssnittet för ditt API
type Server struct {
	e *echo.Echo
}

// NewServer creates a new server instance
func NewServer(logFormat string) *Server {
	fmt.Println("----> Log Format: ", logFormat)

	// Skapa ny Echo-instans
	e := echo.New()

	// Använd din egen logger (MyLogger)
	logger := Logger()
	// Ställ in loggutgången för Echo till att använda din logger
	e.Logger = logger

	logger.Info("logger created")
	logger.Infof("logger format %s", logFormat)

	// Ställ in loggformat baserat på flaggan
	if logFormat == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
		// Lägg till loggning och återställningsmiddleware
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: logger.Output(),
		}))
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `${time_rfc3339} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human}` + "\n",
			Output: logger.Output(),
		}))
	}

	logger.Info("check the output")

	e.Use(middleware.Recover())

	// Skapa serverinstans och hantera routes
	server := &Server{e: e}

	// Statiska filer (CSS, JS, bilder)
	e.Static("/static", "static")

	// Servera Swagger UI-filer
	e.GET("/swagger-ui", func(c echo.Context) error {
		return c.Blob(http.StatusOK, echo.MIMETextHTMLCharsetUTF8, templates.MustGetFile("swagger.html"))
	})

	// Endpoint för att visa OpenAPI-specifikationen
	e.GET("/openapi.yaml", ApiYAMLHandler)

	// Lägg till dina egna routes
	e.GET("/", server.handleHome)
	e.GET("/about", server.handleAbout)

	// Registrera API-handlers
	api.RegisterHandlersWithBaseURL(e, server, "/api")

	return server
}

func setupLogging(logFormat string) *logrus.Logger {
	log := logrus.New()

	switch logFormat {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return log
}

// ApiYAMLHandler hanterar att returnera OpenAPI-specifikationen som YAML
func ApiYAMLHandler(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/x-yaml", api.MustGetFile("api.yaml"))
}

// GetItems implementerar GET /items
func (s *Server) GetItems(ctx echo.Context) error {
	// Skapa objekt med pekarvärden
	id1 := 1
	name1 := "Item 1"
	price1 := float32(10.0)

	id2 := 2
	name2 := "Item 2"
	price2 := float32(20.0)

	items := []api.Item{
		{ID: &id1, Name: &name1, Price: &price1},
		{ID: &id2, Name: &name2, Price: &price2},
	}
	return ctx.JSON(http.StatusOK, items)
}

// CreateItem implementerar POST /items
func (s *Server) CreateItem(ctx echo.Context) error {
	var newItem api.Item
	if err := ctx.Bind(&newItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	// Simulera att objektet fått ett nytt ID
	id := 3
	newItem.ID = &id
	return ctx.JSON(http.StatusCreated, newItem)
}

// GetItemById implementerar GET /items/{id}
func (s *Server) GetItemById(ctx echo.Context, id int) error {
	name := "Item 1"
	price := float32(10.0)
	item := api.Item{ID: &id, Name: &name, Price: &price}
	return ctx.JSON(http.StatusOK, item)
}

// UpdateItemById implementerar PUT /items/{id}
func (s *Server) UpdateItemById(ctx echo.Context, id int) error {
	var updatedItem api.Item
	if err := ctx.Bind(&updatedItem); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	updatedItem.ID = &id
	return ctx.JSON(http.StatusOK, updatedItem)
}

// HeadItemById implementerar HEAD /items/{id}
func (s *Server) HeadItemById(ctx echo.Context, id int) error {
	// Kontrollera om ett objekt existerar
	return ctx.NoContent(http.StatusOK) // Simulera att objektet finns
}

// Start startar servern
func (s *Server) Start(addr string) error {
	return s.e.Start(addr)
}
