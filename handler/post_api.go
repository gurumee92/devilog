package handler

import (
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
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "입력 값이 틀렸습니다.: 바인드 에러 데이터 폼을 살펴보세요."})
	}

	if err := c.Validate(requestDto); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "입력 값이 틀렸습니다.: 유효성 에러 데이터 폼을 살펴보세요."})
	}

	if requestDto.Author == "" {
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "Author는 빈 값이 될 수 없습니다"})
	}

	store := h.postStore
	saved, err := store.Save(&model.Post{
		Title:   requestDto.Title,
		Content: requestDto.Content,
		Author:  requestDto.Author,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.ErrorResponse{Message: "데이터 저장에 실패했습니다.\n" + err.Error()})
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
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "입력 값이 틀렸습니다.: 경로를 살펴보세요."})
	}

	store := h.postStore
	find, err := store.FindByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, &dto.ErrorResponse{Message: "해당 데이터가 존재하지 않습니다."})
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
	requestDto := new(dto.UpdatePostRequestDto)

	if err := c.Bind(requestDto); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "입력 값이 틀렸습니다.: 바인드 에러 데이터 폼을 살펴보세요."})
	}

	if err := c.Validate(requestDto); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "입력 값이 틀렸습니다.: 유효성 에러 데이터 폼을 살펴보세요."})
	}

	pathParam := c.Param("id")
	id, err := strconv.Atoi(pathParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "입력 값이 틀렸습니다.: 경로를 살펴보세요."})
	}

	store := h.postStore
	find, err := store.FindByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, &dto.ErrorResponse{Message: "해당 데이터가 존재하지 않습니다."})
	}

	find.Title = requestDto.Title
	find.Content = requestDto.Content

	updated, err := store.Update(find)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.ErrorResponse{Message: "데이터 저장에 실패했습니다.\n" + err.Error()})
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
