package api

import (
	"btube/serializer"
	"btube/service"

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
			return
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
			return
		}
	}
	c.JSON(200, serializer.ErrResponse(err))
}
