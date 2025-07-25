package books

func (s *Service) GetById(id int) (*Model, error) {
	b := &Model{}
	dbB, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	b.FillFromDB(dbB)
	return b, nil
}
