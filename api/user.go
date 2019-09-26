package api

import (
	"btube/conf"
	"btube/model"
	"btube/serializer"
	"btube/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UserRigister use for user rigistry.
func UserRigister(c *gin.Context) {
	var (
		srv service.UserRigisterService
		err error
	)
	if err = c.ShouldBind(srv); err == nil {
		if user, response := srv.Rigistry(); response != nil {
			c.JSON(200, response)
		} else {
			res := serializer.BuildUserResponse(&user)
			c.JSON(200, res)
		}
		return
	}
	c.JSON(200, serializer.ErrResponse(err))
}

//UserLogin is use for user login and return user massage when success.
func UserLogin(c *gin.Context) {
	var srv service.UserLoginService
	if err := c.ShouldBind(srv); err == nil {
		if user, resp := srv.Login(); resp != nil {
			c.JSON(200, resp)
		} else {
			session := sessions.Default(c)
			session.Clear()
			session.Set("user_id", user.ID)
			session.Save()
			resp := serializer.BuildUserResponse(&user)
			c.JSON(200, resp)
		}
	} else {
		c.JSON(200, serializer.ErrResponse(err))
	}
}

//UserLogout use for user logout.
func UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	c.JSON(200, serializer.Response{
		StatusCode: 0,
		Msg:        "logout successfully.",
	})
}

//CurrentUser get current user.
func CurrentUser(c *gin.Context) *model.User {
	var (
		user *model.User
	)

	session := sessions.Default(c)
	id := session.Get("user_id")
	ID, ok := id.(string)
	if !ok {
		return nil
	}
	err := conf.MySQLConnect.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return nil
	}
	return user
}

//GetUserInformation return the user's information to user.
func GetUserInformation(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(user)
	c.JSON(200, res)
}
