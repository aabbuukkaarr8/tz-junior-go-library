package books

func (r *Repository) Delete(id int) error {
	_, err := r.store.GetConn().Exec(
		`DELETE FROM books WHERE id = $1`, id,
	)
	return err
}
