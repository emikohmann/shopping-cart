package users

import (
	"github.com/emikohmann/shopping-cart/api/src/api/domain/auth"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
	Auth     auth.Auth `json:"auth" gorm:"-"`
}
