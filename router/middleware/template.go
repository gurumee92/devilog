package middleware

import (
	"html/template"
	"io"

	"github.com/gurumee92/devilog/config"
	"github.com/labstack/echo/v4"
)

// Template is ...
type Template struct {
	templates *template.Template
}

// Render is ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplate is
func NewTemplate(c *config.Config) *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob(c.ApplicationPath + "/public/templates/*.html")),
	}
	return t
}
