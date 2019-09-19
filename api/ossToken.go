package api

import (
	"btube/serializer"
	"btube/service"

	"github.com/gin-gonic/gin"
)

//GetOssTokenForVideo get the oss token for user to upload or visid
func GetOssTokenForVideo(c *gin.Context) {
	var srv service.OSSUploadTokenService
	if err := c.ShouldBind(&srv); err == nil {
		resp := srv.GetSignURLForVideo()
		c.JSON(200, resp)
	} else {
		c.JSON(200, serializer.ErrResponse(err))
	}
}

//GetOssTokenForAvatar get the oss token for user to upload or visid
func GetOssTokenForAvatar(c *gin.Context) {
	var srv service.OSSUploadTokenService
	if err := c.ShouldBind(&srv); err == nil {
		resp := srv.GetSignURLForAvatar()
		c.JSON(200, resp)
	} else {
		c.JSON(200, serializer.ErrResponse(err))
	}
}
