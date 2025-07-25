package books

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
