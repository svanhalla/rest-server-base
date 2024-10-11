package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/svanhalla/base-rest-server/internal"

	"github.com/svanhalla/base-rest-server/api"
)

var (
	name      = "base-rest-server"
	version   = "dev"
	buildDate = "undefined"
	user      = "unknown"
)

// YAMLHandler hanterar att returnera OpenAPI-specifikationen som YAML
func YAMLHandler(c echo.Context) error {
	return c.File("api/api.yaml")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Servera Swagger UI-filer
	e.File("/swagger-ui", "templates/swagger.html")

	// Endpoint för att visa OpenAPI-specifikationen
	e.GET("/openapi.yaml", YAMLHandler)

	// Skapa serverinstans och registrera routen med hjälp av oapi-codegen
	server := &internal.Server{}
	api.RegisterHandlersWithBaseURL(e, server, "/api")

	log.Fatal(e.Start(":8080"))
}
