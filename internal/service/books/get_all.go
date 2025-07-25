package books

func (s *Service) GetAll() ([]Model, error) {
	dbBooksPtr, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	books := make([]Model, 0, len(dbBooksPtr))
	for _, dbB := range dbBooksPtr {
		var book Model
		book.FillFromDB(&dbB)
		books = append(books, book)
	}
	return books, nil
}
