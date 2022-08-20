package domain

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID           uint          `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Name         string        `gorm:"column:name;not null;type:varchar(100)" json:"name"`
	Email        string        `gorm:"column:email;unique;not null" json:"email"`
	Password     string        `gorm:"column:password;not null;type:varchar(225)" json:"-"`
	Participants []Participant `gorm:"foreignKey:user_id" json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (User) Migrate(db *gorm.DB) {
	if !db.Migrator().HasTable(User{}.TableName()) {
		if err := db.Migrator().AutoMigrate(User{}); err != nil {
			log.Panicln(fmt.Sprintf(
				"MIGRATE_ERROR(%s): %s",
				User{}.TableName(),
				err.Error(),
			))
		}
	}
}

type UserLoginForm struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"required" msg:"error_invalid_email"`
	Password string `json:"password" form:"password" binding:"required,gte=6,lte=12" validate:"required" msg:"error_invalid_password"`
}

type UserRegisterForm struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"required" msg:"error_invalid_name"`
	Email    string `json:"email" form:"email" binding:"required" validate:"required" msg:"error_invalid_email"`
	Password string `json:"password" form:"password" binding:"required,gte=6,lte=12" validate:"required" msg:"error_invalid_password"`
}

var UserFormErrorMessages = map[string]string{
	"error_invalid_name":     "the name filed is required",
	"error_invalid_email":    "the email filed is required",
	"error_invalid_password": "the password filed is required, and must be (>=6) & (<=12) characters",
}
