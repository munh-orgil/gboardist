package param_handlers

import (
	param_models "gboardist/modules/param/models"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/craftzbay/go_grc/v2/helpers"
	"github.com/gofiber/fiber/v2"
)

type ParamHandler struct{}

func (*ParamHandler) List(c *fiber.Ctx) error {
	key := c.Query("key")
	if res, err := param_models.ParamList(key); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*ParamHandler) Create(c *fiber.Ctx) error {
	data := new(param_models.Param)
	if err := c.BodyParser(data); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if errors := gvalidate.Validate(*data); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	if err := data.Create(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c, data)
}

func (*ParamHandler) Update(c *fiber.Ctx) error {
	data := new(param_models.Param)
	if err := c.BodyParser(data); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}

	if errors := gvalidate.Validate(*data); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	if err := data.Update(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}

func (*ParamHandler) Delete(c *fiber.Ctx) error {
	type Req struct {
		Id uint `json:"id" validate:"required"`
	}
	req := new(Req)
	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if err := gvalidate.Validate(*req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	bp := new(param_models.Param)
	bp.Id = req.Id
	if err := bp.Delete(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}
