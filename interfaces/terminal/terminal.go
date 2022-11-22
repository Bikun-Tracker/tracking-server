package terminal

import (
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	ViewService interface {
		GetTerminalInfo(id string) (dto.GetTerminalInfoResponse, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) GetTerminalInfo(id string) (dto.GetTerminalInfoResponse, error) {
	var (
		response           dto.GetTerminalInfoResponse
		terminal           = &dto.Terminal{}
		allTerminalInRoute = &[]dto.Terminal{}
	)

	err := v.application.TerminalService.GetById(id, terminal)
	if err != nil {
		v.shared.Logger.Errorf("error when finding terminal by id, err: %s", err.Error())
		return response, err
	}

	err = v.application.TerminalService.GetAllByRoute(terminal.Route, allTerminalInRoute)
	if err != nil {
		v.shared.Logger.Errorf("error when finding all terminal by route, err: %s", err.Error())
		return response, err
	}

	response = terminal.ToTerminalInfo(*allTerminalInRoute)

	return response, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
