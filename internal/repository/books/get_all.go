package books

func (r *Repository) GetAll() ([]Model, error) {
	var books []Model
	rows, err := r.store.GetConn().Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		b := Model{}
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.PublishedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil

}
