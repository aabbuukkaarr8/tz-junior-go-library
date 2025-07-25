package books

import (
	repoBooks "tz-junior-go-library/internal/repository/books"
)

type Repo interface {
	Create(m *repoBooks.Model) (*repoBooks.Model, error)
	GetAll() ([]repoBooks.Model, error)
	GetById(id int) (*repoBooks.Model, error)
	Update(m *repoBooks.Model) (*repoBooks.Model, error)
	Delete(id int) error
}
