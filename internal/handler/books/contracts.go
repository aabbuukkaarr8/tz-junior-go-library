package books

import (
	booksService "tz-junior-go-library/internal/service/books"
)

type Service interface {
	Create(input booksService.CreateModel) (*booksService.Model, error)
	GetAll() ([]booksService.Model, error)
	GetById(id int) (*booksService.Model, error)
	Update(book booksService.UpdateBook) (*booksService.Model, error)
	Delete(id int) error
}
