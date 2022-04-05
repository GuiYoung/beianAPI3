package model

import (
	"beianAPI/utils/errmsg"
	"beianAPI/utils/jwt"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"unsafe"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(100);not null" json:"userName" validate:"required,min=3,max=10" label:"用户名"`
	PassWord string `gorm:"type:varchar(500);not null" json:"passWord" validate:"required,min=6,max=12" label:"密码"`
}

// CheckUserExit check user exit
func CheckUserExit(user *User) int {
	data := User{}
	if err := Db.Where("user_name = ?", user.UserName).First(&data).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// create user
func CreateUser(user *User) int {

	var err error
	if err = Db.Where("user_name = ?", user.UserName).First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		user.PassWord = Scrypt(user.PassWord)

		if err = Db.Create(user).Error; err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_USERNAME_USED
}

// find user by name
func FindUser(name string) (user *User, code int) {
	if err := Db.Where("user_name = ?", name).First(&user).Error; err != nil {
		return user, errmsg.ERROR_USERNAME_WRONG
	}
	return user, errmsg.SUCCESS
}

// Scrypt password script
func Scrypt(pwd string) string {
	cost := 10
	dk, _ := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	return *(*string)(unsafe.Pointer(&dk))
}

// check user info
func CheckUser(userName string, password string) (token string, code int) {
	var user *User
	if err := Db.Where("user_name = ?", userName).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return token, errmsg.ERROR_USER_NOT_EXIST
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)); err != nil {
		return token, errmsg.ERROR_PASSWORD_WRONG
	}

	var err error
	if token, err = jwt.GenerateToken(userName); err != nil {
		return "", errmsg.ERROR_GENERATE_TOKEN_FAILED
	}
	return token, errmsg.SUCCESS
}
