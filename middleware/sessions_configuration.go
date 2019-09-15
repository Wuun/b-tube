package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//SessionsConfig is use to config server's seession strategy.
func SessionsConfig(secret string) gin.HandlerFunc {
	cookieStore := cookie.NewStore([]byte(secret))
	cookieStore.Options(sessions.Options{
		Path:     "/",
		MaxAge:   10 * 86400,
		Secure:   true,
		HttpOnly: true,
	})
	return sessions.Sessions("btube-session", cookieStore)
}
