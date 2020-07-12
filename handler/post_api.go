package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gurumee92/devilog/handler/dto"
	"github.com/gurumee92/devilog/model"
	"github.com/labstack/echo/v4"
)

// CreatePost is add post handler
func (h *Handler) CreatePost(c echo.Context) error {
	requestDto := new(dto.CreatePostRequestDto)

	if err := c.Bind(requestDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if requestDto.Author == "" {
		return c.JSON(http.StatusBadRequest, errors.New("Author는 빈 값이 될 수 없습니다"))
	}

	store := h.postStore
	saved, err := store.Save(&model.Post{
		Title:   requestDto.Title,
		Content: requestDto.Content,
		Author:  requestDto.Author,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	responseDto := &dto.PostResponseDto{
		ID:        saved.ID,
		Title:     saved.Title,
		Content:   saved.Content,
		Author:    saved.Author,
		CreatedAt: saved.CreatedAt,
		UpdatedAt: saved.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, responseDto)
}

// GetPost is a post find by ID
func (h *Handler) GetPost(c echo.Context) error {
	pathParam := c.Param("id")
	id, err := strconv.Atoi(pathParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	store := h.postStore
	find, err := store.FindByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	responseDto := &dto.PostResponseDto{
		ID:        find.ID,
		Title:     find.Title,
		Content:   find.Content,
		Author:    find.Author,
		CreatedAt: find.CreatedAt,
		UpdatedAt: find.UpdatedAt,
	}
	return c.JSON(http.StatusOK, responseDto)
}

// UpdatePost is update post findbyID
func (h *Handler) UpdatePost(c echo.Context) error {
	pathParam := c.Param("id")
	id, err := strconv.Atoi(pathParam)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	store := h.postStore
	find, err := store.FindByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	requestDto := new(dto.UpdatePostRequestDto)

	if err := c.Bind(requestDto); err != nil {
		return err
	}

	find.Title = requestDto.Title
	find.Content = requestDto.Content

	updated, err := store.Update(find)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	responseDto := &dto.PostResponseDto{
		ID:        updated.ID,
		Title:     updated.Title,
		Content:   updated.Content,
		Author:    updated.Author,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
	}
	return c.JSON(http.StatusOK, responseDto)
}
