package service

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
)

//UserRigisterService is used to charing user's
//rigisting behaver.
type UserRigisterService struct {
	Nickname        string `json:"nickname" form:"nickname" binding:"required,min=2,max=30"`
	UserName        string `json:"user_name" form:"user_name" binding:"requird,min=2,max=30"`
	Password        string `json:"password" form:"password" binding:"require,min=8,max=30"`
	PasswordConfirm string `josn:"password_confirm" form:"password_confirm" binding:"require,min=8,max=30"`
}

//Check check the user's rigistry information is valid or not.
func (srv *UserRigisterService) Check() *serializer.Response {
	var (
		count int
		user  model.User
	)
	if srv.Password != srv.PasswordConfirm {
		return &serializer.Response{
			StatusCode: 40001,
			Msg:        "passwords you input is not the same.",
		}
	}

	if conf.MySQLConnect.Where("nickname = ?", srv.Nickname).Find(&user).Count(&count); count > 0 {
		return &serializer.Response{
			StatusCode: 40001,
			Msg:        "this nickname has been used by others.",
		}
	}

	count = 0
	if conf.MySQLConnect.Where("user_name = ?").Find(&user).Count(&count); count > 0 {
		return &serializer.Response{
			StatusCode: 40001,
			Msg:        "user name has been used by others.",
		}
	}
	return nil
}

//Rigistry is uesd for user rigistry.
func (srv *UserRigisterService) Rigistry() (user model.User, resp *serializer.Response) {
	user = model.User{
		NickName: srv.Nickname,
		UserName: srv.UserName,
		Status:   model.ActiveUser,
	}

	if response := srv.Check(); response != nil {
		return user, response
	}

	if err := user.SetPassword(srv.Password); err != nil {
		return user, &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to load password.",
		}
	}

	if err := conf.MySQLConnect.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			StatusCode: 40002,
			Msg:        "error when try to registry.",
		}
	}
	return user, nil
}
