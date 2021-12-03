package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitSession() gin.HandlerFunc {
	cookieStore := cookie.NewStore([]byte("bingbing_secret_20210902"))
	cookieStore.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
	})
	return sessions.Sessions("manager_server", cookieStore)
}
