package user_models

import (
	"gboardist/database"
	"gboardist/utils"
	"strings"

	"github.com/craftzbay/go_grc/v2/data"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id             uint            `json:"id" gorm:"primaryKey"`
	Username       string          `json:"username" gorm:"type:varchar(255)"`
	Password       string          `json:"-" gorm:"type:varchar(255)"`
	Email          string          `json:"email" gorm:"type:varchar(255)"`
	PhoneNo        string          `json:"phone_no" gorm:"type:varchar(255)"`
	RegNo          string          `json:"reg_no" gorm:"type:varchar(255)"`
	LastName       string          `json:"last_name" gorm:"type:varchar(150)"`
	FirstName      string          `json:"first_name" gorm:"type:varchar(150)"`
	Gender         uint            `json:"gender"`
	BirthDate      *data.LocalDate `json:"birth_date" gorm:"type:date"`
	CountryCode    string          `json:"country_code" gorm:"type:varchar(10)"`
	OrgId          uint            `json:"org_id"`
	ProfilePicture string          `json:"profile_picture" gorm:"type:varchar(255)"`
	CreatedAt      data.LocalTime  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      data.LocalTime  `json:"updated_at" gorm:"autoUpdateTime"`
}

type ReqUserUpdate struct {
	Id             uint            `json:"id" validate:"required"`
	Username       string          `json:"username"`
	Password       string          `json:"password"`
	PhoneNo        string          `json:"phone_no"`
	RegNo          string          `json:"reg_no"`
	LastName       string          `json:"last_name"`
	FirstName      string          `json:"first_name"`
	Gender         uint            `json:"gender"`
	BirthDate      *data.LocalDate `json:"birth_date"`
	CountryCode    string          `json:"country_code"`
	ProfilePicture string          `json:"profile_picture"`
}

func (p *User) Create() error {
	db := database.DBconn
	return db.Create(&p).Error
}

func (p *User) Update() error {
	db := database.DBconn
	return db.Updates(&p).Error
}

func (p *User) Delete() error {
	db := database.DBconn
	return db.Delete(&p).Error
}

func UserFind(username string) (res *User, err error) {
	db := database.DBconn
	err = db.First(&res, "username = ?", username).Error
	return
}

func UserList(c *fiber.Ctx) (*data.Pagination[User], error) {
	var totalRow int64
	db := database.DBconn
	users := make([]User, 0)
	userName := strings.ToLower(c.Query("username"))

	tx := db.Model(User{})
	if userName != "" {
		tx.Where("username LIKE ?", userName)
	}

	tx.Count(&totalRow)

	p := data.Paginate[User](c, totalRow)

	err := tx.Offset(p.Offset).Limit(p.PageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}
	p.Items = users
	return p, nil
}

func FindUserById(id uint) (res *User, err error) {
	db := database.DBconn
	err = db.Model(&User{}).Take(&res, "id = ?", id).Error
	return
}

func FindUserByUsername(username string) (res *User, err error) {
	db := database.DBconn
	err = db.Model(&User{}).Take(&res, "username = ?", username).Error
	return
}

func FindUserByEmail(email string) (res *User, err error) {
	db := database.DBconn
	err = db.Model(&User{}).Take(&res, "email = ?", email).Error
	return
}

func (u *User) ChangeOrg(orgId uint) error {
	db := database.DBconn

	return db.Model(User{}).Where("id = ?", u.Id).Update("org_id", orgId).Error
}

func GetOrgId(userId uint) (orgId uint, err error) {
	db := database.DBconn
	err = db.Raw("SELECT u.org_id from "+utils.GetTableName("user", "u")+" where u.id=?", userId).Scan(&orgId).Error
	if orgId == 0 {
		err = db.Table(utils.GetTableName("organization_user")).Select("org_id").Where("user_id = ?", userId).Scan(&orgId).Error
	}
	return
}

func GetOrCreateUser(email string) (*User, error) {
	db := database.DBconn
	res, err := FindUserByEmail(email)
	if err == nil {
		return res, nil
	}

	user := User{
		Username: email,
		Email:    email,
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserProfile(userId uint) (res *User, err error) {
	db := database.DBconn
	err = db.First(&res, "id = ?", userId).Error
	return
}
