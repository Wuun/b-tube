package server

import (
	"btube/api"
	"btube/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

//Router return a gin.engine for server's router(router's factory)
func Router() *gin.Engine {
	var router = gin.Default()
	router.Use(middleware.SessionsConfig(os.Getenv("SESSION_SECRET")))
	router.Use(middleware.Cors())
	btube := router.Group("/btube/api/v1", nil)
	{
		btube.GET("ping", nil)
		btube.GET("index", nil)
		btube.POST("registry",api.UserRigister)
		btube.POST("login",api.UserLogin)
		btube.GET("logout",api.UserLogout)
		btube.GET("user_info",api.GetUserInformation)
		btube.POST("avatar_token",api.GetOssTokenForAvatar)
		btube.POST("video_token",api.GetOssTokenForVideo)
		btube.POST("del_video/:id",api.DeleteVideo)
		btube.POST("upload_video",api.UploadVideo)
		btube.GET("list_video",api.ListVideos)
		btube.GET("video_detail/:id",api.VideoDetail)
	}
	return router
}
