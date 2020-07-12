package dto

import "time"

// CreatePostRequestDto is..
type CreatePostRequestDto struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	Author  string `json:"author" form:"author" validate:"required"`
}

// UpdatePostRequestDto is..
type UpdatePostRequestDto struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
}

// PostResponseDto is..
type PostResponseDto struct {
	ID        uint      `json:"id" xml:"id"`
	Title     string    `json:"title" xml:"title"`
	Content   string    `json:"content" xml:"content"`
	Author    string    `json:"author" xml:"author"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt time.Time `json:"updated_at" xml:"updated_at"`
}
