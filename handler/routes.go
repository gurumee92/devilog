package handler

import (
	"github.com/labstack/echo/v4"
)

// Register is routes function
func (h *Handler) Register(e *echo.Echo) {
	e.GET("/hello", h.Hello)
	e.GET("/", h.IndexPage)
	e.GET("/posts/create", h.CreatePostPage)
	e.GET("/posts/update/:id", h.UpdatePostPage)

	postAPIGroup := e.Group("/api/posts")
	postAPIGroup.POST("/", h.CreatePost)
	postAPIGroup.GET("/:id", h.GetPost)
	postAPIGroup.PUT("/:id", h.UpdatePost)
	postAPIGroup.DELETE("/:id", h.DeletePost)

	e.GET("/oauth2/authorization/google", h.GoogleLogin)
	e.GET("/oauth2/authorization/naver", h.NaverLogin)
}
