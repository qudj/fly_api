package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/models"
	"net/http"
)

func ProjectList(c *gin.Context) {
	identity := models.CommonReq{}
	if err := c.ShouldBindQuery(identity); err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
}

func GroupList(c *gin.Context) {
	identity := models.CommonReq{}
	if err := c.ShouldBindQuery(identity); err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
}


func ConfigList(c *gin.Context) {
	identity := models.CommonReq{}
	if err := c.ShouldBindQuery(identity); err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
}


func MiniConfig(c *gin.Context) {
	identity := models.CommonReq{}
	if err := c.ShouldBindQuery(identity); err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"msg":"ok"})
}

