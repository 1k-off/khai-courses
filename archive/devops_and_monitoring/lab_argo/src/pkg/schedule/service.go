package schedule

import "lab-argo-app/pkg/entities"

type Service interface {
	GetSchedule() (*entities.Schedule, error)
	GetDay(day string) (*[]entities.Lesson, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetSchedule() (*entities.Schedule, error) {
	return s.repository.GetSchedule()
}

func (s *service) GetDay(day string) (*[]entities.Lesson, error) {
	return s.repository.GetDay(day)
}
