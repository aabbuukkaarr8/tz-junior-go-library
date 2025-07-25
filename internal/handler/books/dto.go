package books

import (
	"time"
	booksService "tz-junior-go-library/internal/service/books"
)

func (b *UpdateBook) UpdateSrv(id int) booksService.UpdateBook {
	return booksService.UpdateBook{
		ID:     id,
		Title:  b.Title,
		Author: b.Author,
		Pages:  b.Pages,
	}
}

func (m *CreateRequest) ToSrv() booksService.CreateModel {
	return booksService.CreateModel{
		Title:  m.Title,
		Author: m.Author,
		Pages:  m.Pages,
	}
}
func (m *Model) FromSrv(srv *booksService.Model) {
	m.ID = srv.ID
	m.Title = srv.Title
	m.Author = srv.Author
	m.Pages = srv.Pages
	m.PublishedAt = srv.PublishedAt
}

type CreateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type Model struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Pages       int       `json:"pages"`
	PublishedAt time.Time `json:"published-at"`
}

type UpdateBook struct {
	ID     int     `json:"id"`
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Pages  *int    `json:"pages"`
}
