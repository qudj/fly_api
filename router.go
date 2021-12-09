package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/handler"
	"github.com/qudj/fly_api/middlewares"
	"net/http"
)


func initRouter(r *gin.Engine) {
	r.Use(middlewares.SetTraceId)
	r.Use(middlewares.InitSession())

	r.GET("health_check", healthCheck)
	{
		r.Static("doc/swagger_ui", "./swagger_ui")
		r.Static("doc/swagger", "./swagger")
	}

	fcc := r.Group("/fcc")
	//fcc.Use(middlewares.AuthRequired)
	{
		fcc.GET("/project/list/", handler.FccProjectList)
		fcc.GET("/group/list/", handler.FccGroupList)
		fcc.GET("/config/list/", handler.FccConfigList)

		fcc.POST("/project/save/", handler.FccSaveProject)
		fcc.POST("/group/save/", handler.FccSaveGroup)
		fcc.POST("/config/save/", handler.FccSaveConfig)

		fcc.POST("/config/pre_publish/", handler.FccPrePublish)
		fcc.POST("/config/publish/", handler.FccPublish)

		fcc.GET("/config/value", handler.FccGetConfig)
	}

	starling := r.Group("/starling")
	//fcc.Use(middlewares.AuthRequired)
	{
		starling.GET("/project/list/", handler.StarlingProjectList)
		starling.GET("/group/list/", handler.StarlingGroupList)
		starling.GET("/origin/list/", handler.StarlingOriginLgList)
		starling.GET("/translation/list/", handler.StarlingTransLgList)

		starling.POST("/project/save/", handler.StarlingSaveProject)
		starling.POST("/group/save/", handler.StarlingSaveGroup)
		starling.POST("/origin/save/", handler.StarlingSaveOriginLg)
		starling.POST("/translation/save/", handler.StarlingSaveTransLg)

		starling.GET("/translation", handler.StarlingGetTranslation)
	}
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(0, "yes")
	ctx.Status(http.StatusOK)
}