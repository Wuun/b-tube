package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Cors config the cors strategy.
func Cors() gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	c.AllowOrigins = []string{"http://localhost:8080", "https://www.gourouting.com"}
	c.AllowCredentials = true
	return cors.New(c)
}
