package log

import (
	"gboardist/database"
	"strconv"
	"strings"
	"time"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

type Log struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	Path           string         `json:"path" gorm:"type:varchar(100)"`
	Method         string         `json:"method" gorm:"type:varchar(10)"`
	HttpStatusCode int            `json:"http_status_code"`
	Duration       float64        `json:"duration"`
	ReqBody        datatypes.JSON `json:"-"`
	ResBody        datatypes.JSON `json:"-"`
	CreatedAt      data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy      uint           `json:"-"`
}

func RunMigrations() {
	db := database.DBconn
	db.AutoMigrate(Log{})
}

func (s *Log) Save() error {
	db := database.DBconn
	return db.Save(&s).Error
}

func LogList(c *fiber.Ctx) (*data.Pagination[Log], error) {
	var totalRow int64
	db := database.DBconn
	terminals := make([]Log, 0)
	httpStatusCode := strings.ToLower(c.Query("http_status_code"))
	path := strings.ToLower(c.Query("path"))

	tx := db.Model(Log{}).Order("id desc")
	if path != "" {
		tx.Where("path LIKE '%?%'", path)
	}
	if httpStatusCode != "" {
		tx.Where("http_status_code = ?", httpStatusCode)
	}

	tx.Count(&totalRow)

	p := data.Paginate[Log](c, totalRow)

	err := tx.Offset(p.Offset).Limit(p.PageSize).Find(&terminals).Error
	if err != nil {
		return nil, err
	}
	p.Items = terminals
	return p, nil
}

func FiberLogSaver(c *fiber.Ctx, logString []byte) {
	if c.Method() != fiber.MethodGet {
		startTimeStr := c.Get("X-Request-Start-Time")
		startTime, _ := strconv.ParseInt(startTimeStr, 0, 0)
		log := new(Log)
		log.HttpStatusCode = c.Response().StatusCode()
		log.Method = string(c.Request().Header.Method())
		log.Path = string(c.Request().RequestURI())
		log.ReqBody = c.Request().Body()
		log.Duration = float64(time.Now().UnixMicro()-startTime) / 1000000
		if c.Response().StatusCode() != fiber.StatusOK {
			log.ResBody = c.Response().Body()
		}
		log.Save()
	}
}

type ResLog struct {
	Id             uint           `json:"id"`
	Path           string         `json:"path"`
	Method         string         `json:"method"`
	HttpStatusCode int            `json:"http_status_code"`
	ReqBody        datatypes.JSON `json:"req_body"`
	ResBody        datatypes.JSON `json:"res_body"`
	CreatedAt      data.LocalTime `json:"created_at"`
}
