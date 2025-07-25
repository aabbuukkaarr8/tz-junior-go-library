package books

import (
	"time"
)

type Model struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	PublishedAt time.Time `json:"published_at"`
}
