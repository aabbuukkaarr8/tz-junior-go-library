package books

func (s *Service) Update(m UpdateBook) (*Model, error) {
	current, err := s.repo.GetById(m.ID)
	if err != nil {
		return nil, err
	}
	if m.Pages != nil {
		current.Pages = *m.Pages
	}
	if m.Author != nil {
		current.Author = *m.Author
	}
	if m.Title != nil {
		current.Title = *m.Title
	}
	updated, err := s.repo.Update(current)
	if err != nil {
		return nil, err
	}
	fromDb := &Model{}
	fromDb.FillFromDB(updated)

	return fromDb, nil
}
