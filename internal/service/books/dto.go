package books

import (
	"time"
	repoBooks "tz-junior-go-library/internal/repository/books"
)

func (m *Model) FillFromDB(dbm *repoBooks.Model) {
	m.ID = dbm.ID
	m.Title = dbm.Title
	m.Author = dbm.Author
	m.Pages = dbm.Pages
	m.PublishedAt = dbm.PublishedAt
}

type UpdateBook struct {
	ID     int     `json:"id"`
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Pages  *int    `json:"pages"`
}

type CreateModel struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type Model struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	PublishedAt time.Time `json:"published_at"`
}
