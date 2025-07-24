package books

import (
	"time"
	repoBooks "tz-junior-go-library/internal/repository/books"
)

func (s *Service) Create(input CreateModel) (*Model, error) {
	toDB := repoBooks.Model{
		Title:       input.Title,
		Author:      input.Author,
		Pages:       input.Pages,
		PublishedAt: time.Now(),
	}

	created, err := s.repo.Create(&toDB)
	if err != nil {
		return nil, err
	}
	fromDB := Model{
		ID:          created.ID,
		Title:       created.Title,
		Author:      created.Author,
		Pages:       created.Pages,
		PublishedAt: created.PublishedAt,
	}
	return &fromDB, nil
}
