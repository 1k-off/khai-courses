package schedule

import (
	"lab-argo-app/pkg/entities"
)

type Repository interface {
	GetSchedule() (*entities.Schedule, error)
	GetDay(day string) (*[]entities.Lesson, error)
}

type repository struct {
	data entities.Schedule
}

func NewRepository(data entities.Schedule) Repository {
	return &repository{
		data: data,
	}
}

func (r *repository) GetSchedule() (*entities.Schedule, error) {
	return &r.data, nil
}

func (r *repository) GetDay(day string) (*[]entities.Lesson, error) {
	for _, s := range r.data {
		if s.Day == day {
			return &s.Lessons, nil
		}
	}
	return nil, nil
}
