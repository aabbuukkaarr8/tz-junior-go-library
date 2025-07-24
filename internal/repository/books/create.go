package books

import (
	"github.com/google/uuid"
	"time"
)

func (r *Repository) Create(m *Model) (*Model, error) {
	m.PublishedAt = time.Now()
	returnedM := &Model{}
	query := `INSERT INTO books(id, title, author, pages, published_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, title, author, pages, published_at`
	err := r.store.GetConn().QueryRow(
		query,
		uuid.New(),
		m.Title,
		m.Author,
		m.Pages,
		m.PublishedAt,
	).Scan(&returnedM.ID, &returnedM.Title, &returnedM.Author, &returnedM.Pages, &returnedM.PublishedAt)
	if err != nil {
		return nil, err
	}
	return returnedM, nil
}
