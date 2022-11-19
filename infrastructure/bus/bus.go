package bus

import (
	"tracking-server/interfaces"
	"tracking-server/shared"
	"tracking-server/shared/dto"

	"tracking-server/shared/common"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	bus := app.Group("/bus")
	bus.Post("/", c.create)
	bus.Post("/login", c.login)
}

// All godoc
// @Tags Bus
// @Summary Create new bus entry
// @Description Put all mandatory parameter
// @Param CreateBusDto body dto.CreateBusDto true "CreateBus"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateBusResponse
// @Failure 200 {object} dto.CreateBusResponse
// @Router /bus/ [post]
func (c *Controller) create(ctx *fiber.Ctx) error {
	var (
		body     dto.CreateBusDto
		response dto.CreateBusResponse
	)

	err := common.DoCommonRequest(ctx, &body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Shared.Logger.Infof("create bus, data: %s", body)

	response, err = c.Interfaces.BusViewService.CreateBusEntry(body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, response)
}

// All godoc
// @Tags Bus
// @Summary Driver login
// @Description Put all mandatory parameter
// @Param DriverLoginDto body dto.DriverLoginDto true "DriverLoginDto"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.DriverLoginResponse
// @Failure 200 {object} dto.DriverLoginResponse
// @Router /bus/login [post]
func (c *Controller) login(ctx *fiber.Ctx) error {
	var (
		body     dto.DriverLoginDto
		response dto.DriverLoginResponse
	)

	err := common.DoCommonRequest(ctx, &body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	c.Shared.Logger.Infof("login driver, data: %s", body)

	response, err = c.Interfaces.BusViewService.LoginDriver(body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, response)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}
