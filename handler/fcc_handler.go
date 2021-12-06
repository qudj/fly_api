package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/config"
	"github.com/qudj/fly_api/models"
	"github.com/qudj/fly_lib/models/proto/fcc_serv"
	"net/http"
)

func ProjectList(c *gin.Context) {
	// swagger:operation GET /fcc/project/list/ Fcc ProjectListReq
	//
	// Get project list
	//
	// ---
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//             "$ref": "#/definitions/FetchProjectsRet"
	//
	param := &models.ProjectListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.FetchProjectsRequest{
		Limit:       param.Limit,
		Offset:      param.Offset,
		ProjectKey:  param.ProjectKey,
		ProjectName: param.ProjectName,
	}
	res, err := config.FccRpcClient.FetchProjects(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(0, res.Data)
}

func GroupList(c *gin.Context) {
	// swagger:operation GET /fcc/group/list/ Fcc GroupListReq
	//
	// Get group list
	//
	// ---
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//             "$ref": "#/definitions/FetchGroupsRet"
	//
	param := &models.GroupListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.FetchGroupsRequest{
		Limit:  param.Limit,
		Offset: param.Offset,
		ProjectKey: param.ProjectKey,
		GroupKey: param.GroupKey,
		GroupName: param.GroupName,
	}
	res, err := config.FccRpcClient.FetchGroups(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func ConfigList(c *gin.Context) {
	// swagger:operation GET /fcc/config/list/ Fcc ConfigListReq
	//
	// Get config list
	//
	// ---
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//             "$ref": "#/definitions/FetchConfigsRet"
	//
	param := &models.ConfigListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.FetchConfigsRequest{
		Limit:  param.Limit,
		Offset: param.Offset,
		ProjectKey: param.ProjectKey,
		GroupKey: param.GroupKey,
		ConfKey: param.ConfKey,
	}
	res, err := config.FccRpcClient.FetchConfigs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func MiniConfig(c *gin.Context) {
	// swagger:operation GET /fcc/config/value/ Fcc MiniConfigReq
	//
	// Get project list
	//
	// ---
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//             "$ref": "#/definitions/MiniConfig"
	//
	param := &models.MiniConfigReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.FetchMiniConfigRequest{
		ProjectKey: param.ProjectKey,
		GroupKey: param.GroupKey,
		ConfKey: param.ConfKey,
	}
	res, err := config.FccRpcClient.FetchMiniConfig(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}
