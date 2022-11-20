package news

import (
	"tracking-server/interfaces"
	"tracking-server/shared"
	"tracking-server/shared/common"
	"tracking-server/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	news := app.Group("/news")
	news.Post("/", c.create)
	news.Get("/", c.getAll)
}

// All godoc
// @Tags News
// @Summary Create new news
// @Description Put all mandatory parameter
// @Param CreateNewsDto body dto.CreateNewsDto true "CreateNews"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateNewsResponse
// @Failure 200 {object} dto.CreateBusResponse
// @Router /news/ [post]
func (c *Controller) create(ctx *fiber.Ctx) error {
	var (
		body     dto.CreateNewsDto
		response dto.CreateNewsResponse
	)

	err := common.DoCommonRequest(ctx, &body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("create news, data: %s", body)

	response, err = c.Interfaces.NewsViewService.CreateNews(body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, response)
}

// All godoc
// @Tags News
// @Summary Get all news
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetAllNewsResponse
// @Failure 200 {object} dto.GetAllNewsResponse
// @Router /news/ [get]
func (c *Controller) getAll(ctx *fiber.Ctx) error {
	var (
		response dto.GetAllNewsResponse
	)

	response, err := c.Interfaces.NewsViewService.GetAllNews()
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
