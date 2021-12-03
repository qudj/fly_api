package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/config"
	"github.com/qudj/fly_lib/models/fly_conf"
	"net/http"
)

func ProjectList(c *gin.Context) {
	req := &fly_conf.FetchProjectsRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := config.FccRpcClient.FetchProjects(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GroupList(c *gin.Context) {
	req := &fly_conf.FetchGroupsRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := config.FccRpcClient.FetchGroups(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func ConfigList(c *gin.Context) {
	req := &fly_conf.FetchConfigsRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := config.FccRpcClient.FetchConfigs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func MiniConfig(c *gin.Context) {
	req := &fly_conf.FetchMiniConfigRequest{
		ProjectKey: "",
		GroupKey:   "",
		ConfKey:    "",
	}
	res, err := config.FccRpcClient.FetchMiniConfig(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res)
}
