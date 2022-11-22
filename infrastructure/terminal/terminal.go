package terminal

import (
	"tracking-server/interfaces"
	"tracking-server/shared"
	"tracking-server/shared/common"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	news := app.Group("/terminal")
	news.Get("/:id", c.get)
}

// All godoc
// @Tags Terminal
// @Summary Get terminal info
// @Description Put all mandatory parameter
// @Param id path string true "terminal ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetTerminalInfoResponse
// @Failure 200 {object} dto.GetTerminalInfoResponse
// @Router /terminal/{id} [get]
func (c *Controller) get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.Shared.Logger.Infof("get terminal info, data: %s", id)

	res, err := c.Interfaces.TerminalViewsService.GetTerminalInfo(id)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}
