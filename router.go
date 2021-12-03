package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/handler"
	"github.com/qudj/fly_api/middlewares"
	"net/http"
)


func initRouter(r *gin.Engine) {
	// health check
	r.Use(middlewares.InitSession())

	r.GET("health_check", healthCheck)

	fcc := r.Group("/fcc")
	fcc.Use(middlewares.AuthRequired)
	{
		fcc.GET("/project/list", handler.ProjectList)
		fcc.GET("/group/list", handler.GroupList)
		fcc.GET("/config/list", handler.ConfigList)
		fcc.GET("/config", handler.MiniConfig)
	}

}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(0, "yes")
	ctx.Status(http.StatusOK)
}