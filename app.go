package main

import (
	"fmt"
	"gboardist/database"
	"gboardist/global"
	log_module "gboardist/modules/log"
	"gboardist/session"
	"log"
	"os"
	"time"

	"github.com/craftzbay/go_grc/v2/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func main() {
	if err := global.LoadConfig("."); err != nil {
		fmt.Println(err)
		return
	}
	database.InitPostgres()

	cmds := os.Args[1:]
	if helpers.StringInArr("--migrate", cmds) {
		cmds = cmds[1:]
		RunMigrate()
	}

	session.InitJwt(global.Conf.JwtSecretPrvKeyPath, global.Conf.JwtSecretPubKeyPath)
	if len(cmds) == 0 {
		cmds = append(cmds, global.Conf.Port)
	}
	InitServer(cmds[0])
}

func InitServer(port string) {
	zap.L().Info("start project")

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() != fiber.MethodGet {
			c.Request().Header.Add("X-Request-Start-Time", fmt.Sprintf("%d", time.Now().UnixMicro()))
		}

		return c.Next()
	})
	app.Use(logger.New(logger.Config{
		Done: log_module.FiberLogSaver,
	}))

	InitRoutes(app)
	log.Fatal(app.Listen(":" + port))
}
