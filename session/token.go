package session

import (
	"github.com/craftzbay/go_grc/v2/converter"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	IndexUserId  = 0
	IndexLoginId = 1
)

type Token struct {
	jwt.RegisteredClaims
	Grg [8]interface{} `json:"grg"`
}

func GetTokenInfo(c *fiber.Ctx) *Token {
	cs := c.Locals("tokenInfo")
	info, ok := cs.(*Token)
	if !ok {
		return nil
	}
	return info
}

func (tokenInfo *Token) GetUserId() uint {
	if tokenInfo.Grg[IndexUserId] != nil {
		val, _ := converter.InterfaceToUint(tokenInfo.Grg[IndexUserId])
		return val
	}
	return 0
}

func (tokenInfo *Token) SetUserId(userId uint) {
	tokenInfo.Grg[IndexUserId] = userId
}

func (tokenInfo *Token) GetLoginId() uint {
	if tokenInfo.Grg[IndexLoginId] != nil {
		val, _ := converter.InterfaceToUint(tokenInfo.Grg[IndexLoginId])
		return val
	}
	return 0
}

func (tokenInfo *Token) SetLoginId(profileId uint) {
	tokenInfo.Grg[IndexLoginId] = profileId
}
