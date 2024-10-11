package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/svanhalla/base-rest-server/api"
)

// Server implementerar det genererade gränssnittet för ditt API
type Server struct{}

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
