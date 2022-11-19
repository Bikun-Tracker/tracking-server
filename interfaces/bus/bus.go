package bus

import (
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/common"
	"tracking-server/shared/dto"

	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		CreateBusEntry(data dto.CreateBusDto) (dto.CreateBusResponse, error)
		LoginDriver(data dto.DriverLoginDto) (dto.DriverLoginResponse, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateBusEntry(data dto.CreateBusDto) (dto.CreateBusResponse, error) {
	var (
		bus      *dto.Bus
		response dto.CreateBusResponse
	)

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password), bcrypt.DefaultCost,
	)

	if err != nil {
		v.shared.Logger.Errorf("error when encrypting password, err: %s", err.Error())
		return response, err
	}

	bus = &dto.Bus{
		Number:   data.Number,
		Plate:    data.Plate,
		Route:    data.Route,
		Username: data.Username,
		Password: string(encryptedPassword),
	}

	err = v.application.BusService.Create(bus)
	if err != nil {
		v.shared.Logger.Errorf("error when inserting bus to database, err: %s", err.Error())
		return response, err
	}

	response = bus.ToCreateBusResponse()

	return response, nil
}

func (v *viewService) LoginDriver(data dto.DriverLoginDto) (dto.DriverLoginResponse, error) {
	var (
		bus      = &dto.Bus{}
		response dto.DriverLoginResponse
	)

	err := v.application.BusService.FindByUsername(data.Username, bus)
	if err != nil {
		v.shared.Logger.Errorf("error when finding bus by username, err: %s", err.Error())
		return response, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(bus.Password),
		[]byte(data.Password),
	)
	if err != nil {
		v.shared.Logger.Errorf("wrong password, err: %s", err.Error())
		return response, err
	}

	token, err := common.NewJWT(bus.Username, v.shared.Env)
	if err != nil {
		v.shared.Logger.Errorf("error when creating jwt, err: %s", err.Error())
		return response, err
	}

	response = bus.ToDriverLoginResponse(token)

	return response, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
