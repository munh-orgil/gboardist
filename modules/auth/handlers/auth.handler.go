package auth_handlers

// import (
// 	auth_models "gboardist/modules/auth/models"
// 	user_models "gboardist/modules/user/models"
// 	"gboardist/session"
// 	"gboardist/utils"

// 	"github.com/craftzbay/go_grc/v2/gvalidate"
// 	"github.com/craftzbay/go_grc/v2/helpers"
// 	"github.com/gofiber/fiber/v2"
// )

// type AuthHandler struct{}

// func (*AuthHandler) Login(c *fiber.Ctx) error {
// 	req := new(auth_models.ReqLogin)

// 	if err := c.BodyParser(req); err != nil {
// 		return helpers.ResponseBadRequest(c, err.Error())
// 	}
// 	if err := gvalidate.Validate(*req); err != nil {
// 		return helpers.ResponseBadRequest(c, err.Error())
// 	}

// 	user, err := user_models.FindUserByUsername(req.Username)
// 	if err != nil {
// 		return helpers.Response(c, err.Error())
// 	}

// 	if user.Password != helpers.GeneratePassword(req.Password) {
// 		return helpers.ResponseBadRequest(c, "Хэрэглэгчийн мэдээлэл буруу байна")
// 	}

// 	tokenInfo := new(session.Token)
// 	tokenInfo.SetUserId(user.Id)

// 	token, err := session.GetToken(tokenInfo)
// 	if err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}
// 	return helpers.Response(c, session.ResToken{Token: token})
// }

// func (*AuthHandler) Register(c *fiber.Ctx) error {
// 	req := new(auth_models.ReqRegister)

// 	if err := c.BodyParser(req); err != nil {
// 		return helpers.ResponseBadRequest(c, err.Error())
// 	}
// 	if errors := gvalidate.Validate(*req); errors != nil {
// 		return helpers.ResponseBadRequest(c, errors.Error())
// 	}

// 	user := user_models.User{
// 		Username:  req.Email,
// 		Password:  helpers.GeneratePassword(req.Password),
// 		Email:     req.Email,
// 		BirthDate: nil,
// 	}

// 	if err := user.Create(); err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}
// 	return helpers.Response(c)
// }

// func (*AuthHandler) ForgotPassword(c *fiber.Ctx) error {
// 	req := new(auth_models.ReqRegister)

// 	if err := c.BodyParser(req); err != nil {
// 		return helpers.ResponseBadRequest(c, err.Error())
// 	}
// 	if errors := gvalidate.Validate(*req); errors != nil {
// 		return helpers.ResponseBadRequest(c, errors.Error())
// 	}
// 	if err := otp.CheckOtp(req.Email, req.Otp); err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}

// 	user, err := user_models.FindUserByEmail(req.Email)
// 	if err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}
// 	user.Username = req.Email
// 	user.Password = helpers.GeneratePassword(req.Password)

// 	if err := user.Update(); err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}
// 	return helpers.Response(c)
// }

// func (*AuthHandler) ChangePassword(c *fiber.Ctx) error {
// 	req := new(auth_models.ReqChange)

// 	if err := c.BodyParser(req); err != nil {
// 		return helpers.ResponseBadRequest(c, err.Error())
// 	}
// 	if errors := gvalidate.Validate(*req); errors != nil {
// 		return helpers.ResponseBadRequest(c, errors.Error())
// 	}

// 	user, err := user_models.FindUserById(session.GetTokenInfo(c).GetUserId())
// 	if err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}

// 	if user.Password != helpers.GeneratePassword(req.Old) {
// 		return helpers.ResponseBadRequest(c, "Нууц үг буруу")
// 	}
// 	user.Password = helpers.GeneratePassword(req.New)

// 	if err := user.Update(); err != nil {
// 		return helpers.ResponseErr(c, err.Error())
// 	}
// 	return helpers.Response(c)
// }

// func (*AuthHandler) GetOtp(c *fiber.Ctx) error {
// 	email := c.Query("email")
// 	requestType := c.Query("request_type")
// 	if email == "" {
// 		return helpers.ResponseBadRequest(c, "Имэйл шаардлагатай")
// 	}

// 	if !gvalidate.IsEmail(email) {
// 		return helpers.ResponseBadRequest(c, "Имэйл биш байна")
// 	}

// 	if requestType == "forgot" && !utils.CheckExists("user", []string{"email"}, []interface{}{email}) {
// 		return helpers.ResponseBadRequest(c, "Имэйл бүртгэлгүй байна")
// 	}

// 	if requestType != "forgot" && utils.CheckExists("user", []string{"email"}, []interface{}{email}) {
// 		return helpers.ResponseBadRequest(c, "Имэйл бүртгэлтэй байна")
// 	}

// 	if err := otp.SendOtp(email); err != nil {
// 		if err.Error() == "MESSAGE_WAIT_SECOND_ERROR" {
// 			return helpers.Response(c, "OTP аль хэдийн илгээгдсэн байна")
// 		}
// 		return helpers.ResponseBadRequest(c, err.Error())
// 	}

// 	return helpers.Response(c, "OTP код илгээгдлээ")
// }
