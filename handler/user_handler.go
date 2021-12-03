package handler


import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/config"
	"net/http"
)

type LoginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	body := LoginRequestBody{
		Username: "qudongjie",
		Password: "qudongjie_123456",
	}

	if pass, ok := config.Users[body.Username]; ok {
		if pass != body.Password {
			c.JSON(400, "")
			return
		}
	} else {
		c.JSON(400, "")
		return
	}

	// Save the username in the session
	session.Set(config.UserKey, body.Username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(400, "")
		return
	}

	c.JSON(http.StatusOK, "")
}


func Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(config.UserKey)
	c.JSON(http.StatusOK, gin.H{"username": user})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(config.UserKey)
	if user == nil {
		c.JSON(400, "")
		return
	}
	session.Delete(config.UserKey)
	if err := session.Save(); err != nil {
		c.JSON(400, "")
		return
	}
	c.JSON(http.StatusOK, "")
}