package books

import (
	"github.com/google/uuid"
	"time"
)

type CreateModel struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type Model struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	PublishedAt time.Time `json:"published_at"`
}
