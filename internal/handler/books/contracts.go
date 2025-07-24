package books

import (
	booksService "tz-junior-go-library/internal/service/books"
)

type Service interface {
	Create(input booksService.CreateModel) (*booksService.Model, error)
}
