package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/config"
	"github.com/qudj/fly_lib/models/proto/fcc_serv"
	"net/http"
)

func ProjectList(c *gin.Context) {
	req := &fcc_serv.FetchProjectsRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := config.FccRpcClient.FetchProjects(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func GroupList(c *gin.Context) {
	req := &fcc_serv.FetchGroupsRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := config.FccRpcClient.FetchGroups(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func ConfigList(c *gin.Context) {
	req := &fcc_serv.FetchConfigsRequest{
		Limit:  10,
		Offset: 0,
	}
	res, err := config.FccRpcClient.FetchConfigs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func MiniConfig(c *gin.Context) {
	req := &fcc_serv.FetchMiniConfigRequest{
		ProjectKey: "",
		GroupKey:   "",
		ConfKey:    "",
	}
	res, err := config.FccRpcClient.FetchMiniConfig(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}
