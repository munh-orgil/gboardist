package user

import (
	"gboardist/database"
	user_models "gboardist/modules/user/models"
)

func RunMigrations() {
	db := database.DBconn
	db.AutoMigrate(user_models.User{})
}
