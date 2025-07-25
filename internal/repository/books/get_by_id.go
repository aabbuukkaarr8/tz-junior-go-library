package books

func (r *Repository) GetById(id int) (*Model, error) {
	var b Model
	err := r.store.GetConn().QueryRow(`SELECT * FROM books WHERE id = $1`, id).Scan(
		&b.ID, &b.Title, &b.Author, &b.Pages, &b.PublishedAt)
	if err != nil {
		return nil, nil
	}
	return &b, nil
}
