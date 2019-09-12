package server

import "github.com/gin-gonic/gin"

//Router return a gin.engine for server's router(router's factory)
func Router() *gin.Engine {
	router := gin.Default()
	btube := router.Group("/butube/api/v1", nil)
	{
		btube.GET("ping", nil)
		btube.GET("index", nil)
	}
	return router
}
