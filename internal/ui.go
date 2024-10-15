package internal

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

// Load templates
func loadTemplate(file string) *template.Template {
	// return template.Must(template.ParseFS(templates.GetFS(), "layout.html.tmpl", file))
	return template.Must(template.ParseFiles("templates/layout.html.tmpl", file))
}

// Handlers
func (s *Server) handleHome(c echo.Context) error {
	tmpl := loadTemplate("templates/index.html.tmpl")
	data := map[string]interface{}{
		"Title": "Home",
	}
	return tmpl.ExecuteTemplate(c.Response().Writer, "layout", data)
}

func (s *Server) handleAbout(c echo.Context) error {
	tmpl := loadTemplate("templates/about.html.tmpl")
	data := map[string]interface{}{
		"Title": "About Us",
	}
	return tmpl.ExecuteTemplate(c.Response().Writer, "layout", data)
}
