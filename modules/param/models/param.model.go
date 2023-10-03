package param_models

import (
	"gboardist/database"

	"github.com/craftzbay/go_grc/v2/data"
	"gorm.io/datatypes"
)

type Param struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Key       string         `json:"key" gorm:"type:varchar(255)"`
	Value     datatypes.JSON `json:"value"`
	CreatedAt data.LocalTime `json:"created_at" gorm:"autoCreateTime"`
}

func ParamList(key string) (res []Param, err error) {
	db := database.DBconn

	tx := db.Model(Param{})
	if key != "" {
		tx.Where("key = ?", key)
	}

	err = tx.Find(&res).Error
	return
}

func (p *Param) Create() error {
	db := database.DBconn
	return db.Create(&p).Error
}

func (p *Param) Update() error {
	db := database.DBconn
	return db.Where("id = ?", p.Id).Updates(&p).Error
}

func (p *Param) Delete() error {
	db := database.DBconn
	return db.Delete(&p).Error
}
