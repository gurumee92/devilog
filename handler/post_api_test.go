package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gurumee92/devilog/handler/dto"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostSuccess(t *testing.T) {
	requestData := dto.CreatePostRequestDto{
		Title:   "test title",
		Content: "test content",
		Author:  "test author",
	}
	b, err := json.Marshal(requestData)
	assert.NoError(t, err)

	req := httptest.NewRequest(echo.POST, "/api/posts/", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.CreatePost(c)) {
		var resp dto.PostResponseDto
		assert.Equal(t, http.StatusCreated, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Less(t, uint(1), resp.ID)
		assert.Equal(t, requestData.Title, resp.Title)
		assert.Equal(t, requestData.Content, resp.Content)
		assert.Equal(t, requestData.Author, resp.Author)
		assert.Contains(t, rec.Body.String(), "created_at")
		assert.Contains(t, rec.Body.String(), "updated_at")
	}
}

func TestCreatePostFailedBindError(t *testing.T) {
	requestData := `{ "test": "test" }`
	req := httptest.NewRequest(echo.POST, "/api/posts/", strings.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.CreatePost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreatePostFailedEmptyAuthor(t *testing.T) {
	requestData := dto.CreatePostRequestDto{
		Title:   "test title",
		Content: "test content",
	}
	b, err := json.Marshal(requestData)
	assert.NoError(t, err)

	req := httptest.NewRequest(echo.POST, "/api/posts/", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.CreatePost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetPostSuccess(t *testing.T) {
	id := 1
	test := "test" + strconv.Itoa(id)

	req := httptest.NewRequest(echo.GET, "/api/posts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	if assert.NoError(t, h.GetPost(c)) {
		var resp dto.PostResponseDto
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, uint(id), resp.ID)
		assert.Equal(t, test, resp.Title)
		assert.Equal(t, test, resp.Content)
		assert.Equal(t, test, resp.Author)
		assert.Contains(t, rec.Body.String(), "created_at")
		assert.Contains(t, rec.Body.String(), "updated_at")
	}
}

func TestGetPostFailedPathVarialbeTypeError(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/posts/star", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.GetPost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetPostFailedNotFound(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/posts/10", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(10))

	if assert.NoError(t, h.GetPost(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestUpdatePostSuccess(t *testing.T) {
	id := 1
	requestData := dto.UpdatePostRequestDto{
		Title:   "test update title",
		Content: "test update content",
	}
	b, err := json.Marshal(requestData)
	assert.NoError(t, err)

	req := httptest.NewRequest(echo.PUT, "/api/posts/1", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	if assert.NoError(t, h.UpdatePost(c)) {
		var resp dto.PostResponseDto
		assert.Equal(t, http.StatusOK, rec.Code)
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, uint(id), resp.ID)
		assert.Equal(t, requestData.Title, resp.Title)
		assert.Equal(t, requestData.Content, resp.Content)
		assert.Contains(t, rec.Body.String(), "author")
		assert.Contains(t, rec.Body.String(), "created_at")
		assert.Contains(t, rec.Body.String(), "updated_at")
	}
}

func TestUpdatePostFailedNotFound(t *testing.T) {
	id := 10
	requestData := dto.UpdatePostRequestDto{
		Title:   "test update title",
		Content: "test update content",
	}
	b, err := json.Marshal(requestData)
	assert.NoError(t, err)

	req := httptest.NewRequest(echo.PUT, "/api/posts/10", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(id))

	if assert.NoError(t, h.GetPost(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestUpdatePostFailedPathVariableTypeError(t *testing.T) {
	id := "start"
	requestData := dto.UpdatePostRequestDto{
		Title:   "test update title",
		Content: "test update content",
	}
	b, err := json.Marshal(requestData)
	assert.NoError(t, err)

	req := httptest.NewRequest(echo.PUT, "/api/posts/10", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	if assert.NoError(t, h.GetPost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdatePostFailedBindError(t *testing.T) {
	requestData := `{ "test": "test" }`

	req := httptest.NewRequest(echo.PUT, "/api/posts/1", strings.NewReader(requestData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.GetPost(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
