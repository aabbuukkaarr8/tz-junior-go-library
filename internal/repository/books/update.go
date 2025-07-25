package books

func (r *Repository) Update(m *Model) (*Model, error) {
	updated := Model{}
	query := `UPDATE books SET title = $1, author = $2, pages = $3 WHERE id = $4
RETURNING id, title, author, pages, published_at`
	err := r.store.GetConn().QueryRow(query,
		m.Title,
		m.Author,
		m.Pages,
		m.ID).Scan(&updated.ID, &updated.Title, &updated.Author, &updated.Pages, &updated.PublishedAt)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}
