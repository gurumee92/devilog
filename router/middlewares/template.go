package middlewares

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template is
type Template struct {
	templates *template.Template
}

// Render is
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplate is ..
func NewTemplate() *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	return t
}
