package bus

import (
	"time"
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/common"
	"tracking-server/shared/dto"

	"github.com/gofiber/websocket/v2"
	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		CreateBusEntry(data dto.CreateBusDto) (dto.CreateBusResponse, error)
		LoginDriver(data dto.DriverLoginDto) (dto.DriverLoginResponse, error)
		DeleteBus(id string) error
		EditBus(data dto.EditBusDto, id string, token string) (dto.EditBusResponse, error)
		TrackBusLocation(token string, c *websocket.Conn) error
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

func (v *viewService) DeleteBus(id string) error {
	err := v.application.BusService.Delete(id)
	if err != nil {
		v.shared.Logger.Errorf("error when deleteing bus, err: %s", err.Error())
		return err
	}
	return nil
}

func (v *viewService) EditBus(data dto.EditBusDto, id string, token string) (dto.EditBusResponse, error) {
	var (
		bus      = &dto.Bus{}
		response dto.EditBusResponse
	)

	username, err := common.ExtractTokenData(token, v.shared.Env)
	if err != nil {
		v.shared.Logger.Errorf("error when extract jwt, err: %s", err.Error())
	}

	err = v.application.BusService.FindByUsername(username, bus)
	if err != nil {
		v.shared.Logger.Errorf("error when checking username, err: %s", err.Error())
		return response, err
	}

	err = v.application.BusService.FindById(id, bus)
	if err != nil {
		v.shared.Logger.Errorf("error when finding bus, err: %s", err.Error())
		return response, err
	}

	bus.FillBusEdit(data)

	err = v.application.BusService.Save(bus)
	if err != nil {
		v.shared.Logger.Errorf("error when saving update, err: %s", err.Error())
		return response, err
	}

	response = bus.ToEditBusResponnse()

	return response, nil
}

func (v *viewService) TrackBusLocation(token string, c *websocket.Conn) error {
	var (
		data     = dto.BusLocationMessage{}
		bus      = dto.Bus{}
		location = dto.BusLocation{}
	)

	username, err := common.ExtractTokenData(token, v.shared.Env)
	if err != nil {
		v.shared.Logger.Errorf("error when parsing jwt, err: %s", err.Error())
		c.Close()
		return err
	}

	err = v.application.BusService.FindByUsername(username, &bus)
	if err != nil {
		c.Close()
		return err
	}

	if err := c.ReadJSON(&data); err != nil {
		v.shared.Logger.Errorf("error when receiving websocket message, err: %s", err.Error())
		c.Close()
		return err
	}

	location.BusID = bus.ID
	location.Lat = data.Lat
	location.Long = data.Long
	location.Timestamp = time.Now()

	go func() {
		v.application.BusService.InsertBusLocation(&location)
		v.shared.Logger.Infof("insert bus location, data: %s", location)
	}()

	if err := c.WriteJSON(data); err != nil {
		v.shared.Logger.Errorf("error when writing message, err: %s", err.Error())
		c.Close()
		return err
	}

	return nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
