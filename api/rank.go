package api

import (
	"btube/service"
	"github.com/gin-gonic/gin"
)

func Rank(c *gin.Context){
	var srv service.DailyRankService
	c.JSON(200,srv.Rank())
}