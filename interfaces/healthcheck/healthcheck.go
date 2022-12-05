package healthcheck

import (
	"tracking-server/application"
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	ViewService interface {
		SystemHealthcheck() (dto.HCStatus, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

/**
 * Get status for http resolver and database
 */
func (v *viewService) SystemHealthcheck() (dto.HCStatus, error) {
	status := make([]dto.Status, 0)

	httpStatus := v.application.HealthcheckService.HttpHealthcheck(v.shared.Http)
	status = append(status, httpStatus)

	dbStatus := v.application.HealthcheckService.DatabaseHealthcheck(v.shared.DB)
	status = append(status, dbStatus)

	return dto.HCStatus{
		Status: status,
	}, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
