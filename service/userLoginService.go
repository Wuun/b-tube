package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//UserLoginService is a service used for user login.
type UserLoginService struct {
	Nickname string `json:"nickname" form:"nickname" binding:"required,min=3,max=30"`
	Password string `json:"passwoed" form:"password" binding:"required,min=8,max=40"`
}

//Login is use for user login and return user's information when success.
func (srv *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User
	if err := conf.MySQLConnect.Where("nickname = ?", srv.Nickname).First(&user).Error; err != nil {
		return user, &serializer.Response{
			StatusCode: 40001,
			Msg:        "wrong nickname or password.",
		}
	}
	if !user.CheckPassword(srv.Password) {
		return user, &serializer.Response{
			StatusCode: 40001,
			Msg:        "wrong nickname or password.",
		}
	}
	return user, nil
}
