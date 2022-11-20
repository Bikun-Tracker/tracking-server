package news

import (
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	Service interface {
		Create(data *dto.News) error
	}
	service struct {
		shared shared.Holder
	}
)

func (s *service) Create(data *dto.News) error {
	err := s.shared.DB.Create(data).Error
	return err
}

func NewNewsService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}
