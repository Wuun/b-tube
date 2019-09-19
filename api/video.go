package api

import (
	"btube/serializer"
	"btube/service"

	"github.com/gin-gonic/gin"
)

//UploadVideo is api used for upload video.
func UploadVideo(c *gin.Context) {
	var upService *service.UploadVideoService
	if err := c.ShouldBind(&upService); err == nil {
		c.JSON(200, upService.Upload())
	} else {
		c.JSON(200, serializer.ErrResponse(err))
	}
}
