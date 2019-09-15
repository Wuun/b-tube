package api

import (
	"btube/serializer"

	"github.com/gin-gonic/gin"
)

//Ping is ues to test the server is down or not.
func Ping(c *gin.Context) {
	c.JSON(
		200,
		serializer.Response{
			StatusCode: 0,
			Msg:        "Pong!",
		})
}
