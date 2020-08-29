package service

type Service struct {
	ss StoriesService
}

func (s *Service) StoriesService() StoriesService {
	return s.ss
}

func NewService(ss StoriesService) *Service {
	return &Service{
		ss: ss,
	}
}
