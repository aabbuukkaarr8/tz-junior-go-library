package books

import "tz-junior-go-library/internal/store"

type Repository struct {
	store *store.Store
}

func NewRepository(store *store.Store) *Repository {
	return &Repository{
		store: store,
	}
}
