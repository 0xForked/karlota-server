package domain

type User struct {
	ID       int64  `gorm:"column:id;primarykey" sql:"index" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email;unique" json:"email" binding:"required"`
	Password string `gorm:"column:password" json:"-" binding:"required"`
}

func (User) TableName() string {
	return "users"
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

var UserFromErrorMessages = map[string]string{
	"error_invalid_name":     "the name filed is required",
	"error_invalid_email":    "the email filed is required",
	"error_invalid_password": "the password filed is required, and must be (>=6) & (<=12) characters",
}
