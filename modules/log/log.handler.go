package log

import (
	"gboardist/database"

	"github.com/craftzbay/go_grc/v2/helpers"
	"github.com/gofiber/fiber/v2"
)

type LogHandler struct{}

func (*LogHandler) Test(c *fiber.Ctx) error {
	return helpers.Response(c, "test done")
}

func (*LogHandler) List(c *fiber.Ctx) error {
	var (
		res interface{}
		err error
	)

	if res, err = LogList(c); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}

	return helpers.Response(c, (res))
}

func (*LogHandler) Find(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	var res ResLog

	db := database.DBconn
	if err := db.Model(&Log{}).Take(&res, id).Error; err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c, res)
}
