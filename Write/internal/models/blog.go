package models

import "time"

type Blog struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" validate:"required,min=3,max=200"`
	Description string    `json:"description" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Metadata    Metadata  `json:"metadata"`
}

type Metadata struct {
	Likes    int    `json:"likes"`
	Category string `json:"category"`
}


type ElasticBlog struct{
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Tags        []string  `json:"tags"`
	Category    string    `json:"category"`
	Likes       int       `json:"likes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}