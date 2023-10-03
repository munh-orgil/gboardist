package param

import (
	"gboardist/database"
	param_models "gboardist/modules/param/models"
)

func RunMigrations() {
	db := database.DBconn
	db.AutoMigrate(param_models.Param{})
}
