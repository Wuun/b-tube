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

//ListVideos is use to list all video the server has.
func ListVideos(c *gin.Context) {
	srv := service.ListVideoService{}
	if err := c.ShouldBind(&srv); err == nil {
		c.JSON(200, srv.List())
	} else {
		c.JSON(200, serializer.ErrResponse(err))
	}
}

//VideoDetail list the video detail.
func VideoDetail(c *gin.Context) {
	var srv service.VideoDetailService
	c.JSON(200, srv.Show(c.Param("id")))
}

//DeleteVideo delete the video.
func DeleteVideo(c *gin.Context) {
	var srv service.DeleteVideoService
	id := c.Param("id")
	c.JSON(200, srv.Delete(id))
}
