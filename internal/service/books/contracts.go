package books

import (
	repoBooks "tz-junior-go-library/internal/repository/books"
)

type Repo interface {
	Create(m *repoBooks.Model) (*repoBooks.Model, error)
}
