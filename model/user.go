package model

import (
	"btube/conf"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User is thw model of user.
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

const (
	PasswordCost = 12
	ActiveUser   = "active_user"
	InactiveUser = "inactive_user"
	SuspendUser  = "suspend_user"
)

// GetUser get the user information.
func GetUser(ID interface{}) (User, error) {
	var user User
	result := conf.MySQLConnect.First(&user, ID)
	return user, result.Error
}

// SetPassword set the password of user.
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword validate the user's password.
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
