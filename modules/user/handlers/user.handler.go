package user_handlers

import (
	user_models "gboardist/modules/user/models"
	"gboardist/session"

	"github.com/craftzbay/go_grc/v2/gvalidate"
	"github.com/craftzbay/go_grc/v2/helpers"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func (*UserHandler) Find(c *fiber.Ctx) error {
	username := c.Query("search_text")

	if username == "" {
		return helpers.ResponseBadRequest(c, "Нэвтрэх нэр оруулна уу")
	}
	if res, err := user_models.UserFind(username); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*UserHandler) List(c *fiber.Ctx) error {
	if res, err := user_models.UserList(c); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, res)
	}
}

func (*UserHandler) Profile(c *fiber.Ctx) error {
	userId := session.GetTokenInfo(c).GetUserId()

	if profile, err := user_models.GetUserProfile(userId); err != nil {
		return helpers.ResponseErr(c, err.Error())
	} else {
		return helpers.Response(c, profile)
	}
}

func (*UserHandler) Update(c *fiber.Ctx) error {
	req := new(user_models.ReqUserUpdate)

	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if errors := gvalidate.Validate(*req); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	user := user_models.User{
		Id:          req.Id,
		Username:    req.Username,
		PhoneNo:     req.PhoneNo,
		RegNo:       req.RegNo,
		LastName:    req.LastName,
		FirstName:   req.FirstName,
		Gender:      req.Gender,
		BirthDate:   req.BirthDate,
		CountryCode: req.CountryCode,
	}
	if req.Password != "" {
		user.Password = helpers.GeneratePassword(req.Password)
	}

	if err := user.Update(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}

func (*UserHandler) Delete(c *fiber.Ctx) error {
	type Req struct {
		Id       uint   `json:"id" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	req := new(Req)
	if err := c.BodyParser(req); err != nil {
		return helpers.ResponseBadRequest(c, err.Error())
	}
	if errors := gvalidate.Validate(*req); errors != nil {
		return helpers.ResponseBadRequest(c, errors.Error())
	}

	deletedUser, err := user_models.FindUserById(req.Id)
	if err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	if deletedUser.Password != helpers.GeneratePassword(req.Password) {
		return helpers.ResponseBadRequest(c, "Нууц үг буруу байна")
	}

	if err := deletedUser.Delete(); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}

	return helpers.Response(c)
}

func (*UserHandler) ChangeOrg(c *fiber.Ctx) error {
	orgId := uint(c.QueryInt("org_id"))
	if orgId == 0 {
		return helpers.ResponseErr(c, "org_id is required")
	}

	user := user_models.User{
		Id: session.GetTokenInfo(c).GetUserId(),
	}

	if err := user.ChangeOrg(orgId); err != nil {
		return helpers.ResponseErr(c, err.Error())
	}
	return helpers.Response(c)
}
