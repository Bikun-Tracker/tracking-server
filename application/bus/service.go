package bus

import (
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	Service interface {
		Create(data *dto.Bus) error
		FindByUsername(username string, bus *dto.Bus) error
		Delete(id string) error
	}
	service struct {
		shared shared.Holder
	}
)

func (s *service) Create(data *dto.Bus) error {
	err := s.shared.DB.Create(data).Error
	return err
}

func (s *service) FindByUsername(username string, bus *dto.Bus) error {
	err := s.shared.DB.Where("username = ?", username).First(bus).Error
	return err
}

func (s *service) Delete(id string) error {
	err := s.shared.DB.Delete(&dto.Bus{}, id).Error
	return err
}

func NewBusService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}
