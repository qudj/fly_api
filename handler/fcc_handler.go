package handler

import (
	"fmt"
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
		Limit:  param.Limit,
		Offset: param.Offset,
		Filter: map[string]string{},
	}
	if param.ProjectKey != "" {
		req.Filter["project_key"] = param.ProjectKey
	}
	if param.ProjectName != "" {
		req.Filter["project_name"] = param.ProjectName
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
		Limit:      param.Limit,
		Offset:     param.Offset,
		ProjectKey: param.ProjectKey,
		Filter:     map[string]string{},
	}
	if param.GroupKey != "" {
		req.Filter["group_key"] = param.GroupKey
	}
	if param.GroupName != "" {
		req.Filter["group_name"] = param.GroupName
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
		Limit:      param.Limit,
		Offset:     param.Offset,
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		Filter:     map[string]string{},
	}
	if param.ConfKey != "" {
		req.Filter["conf_key"] = param.ConfKey
	}
	res, err := config.FccRpcClient.FetchConfigs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func SaveProject(c *gin.Context) {
	// swagger:operation POST /fcc/project/save/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/SaveProjectReq"
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
	param := &models.SaveProjectReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.SaveProjectRequest{
		Project: &fcc_serv.Project{
			ProjectKey:  param.ProjectKey,
			ProjectName: param.ProjectName,
			Description: param.Description,
			Status:      param.Status,
		},
		OpId: "qudongjie",
	}
	res, err := config.FccRpcClient.SaveProject(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res == nil || res.BaseRet == nil || res.BaseRet.Code != 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("res=%+v", res))
		return
	}
	c.JSON(http.StatusOK, "")
}

func SaveGroup(c *gin.Context) {
	// swagger:operation POST /fcc/group/save/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/SaveGroupReq"
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//
	param := &models.SaveGroupReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.SaveGroupRequest{
		Group: &fcc_serv.Group{
			ProjectKey:  param.ProjectKey,
			GroupKey:    param.GroupKey,
			GroupName:   param.GroupName,
			Description: param.Description,
			Status:      param.Status,
		},
		OpId: "qudongjie",
	}
	res, err := config.FccRpcClient.SaveGroup(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res == nil || res.BaseRet == nil || res.BaseRet.Code != 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("res=%+v", res))
		return
	}
	c.JSON(http.StatusOK, "")
}

func SaveConfig(c *gin.Context) {
	// swagger:operation POST /fcc/config/save/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/SaveConfigReq"
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//
	param := &models.SaveConfigReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.SaveConfigRequest{
		Config: &fcc_serv.Config{
			ProjectKey:  param.ProjectKey,
			GroupKey:    param.GroupKey,
			ConfKey:     param.ConfKey,
			Description: param.Description,
			Status:      param.Status,
		},
		OpId: "qudongjie",
	}
	res, err := config.FccRpcClient.SaveConfig(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res == nil || res.BaseRet == nil || res.BaseRet.Code != 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("res=%+v", res))
		return
	}
	c.JSON(http.StatusOK, "")
}

func PrePublish(c *gin.Context) {
	// swagger:operation POST /fcc/config/pre_publish/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/PrePublishReq"
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
	param := &models.PrePublishReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.PrePublishRequest{
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		ConfKey:    param.ConfKey,
		PreValue:   param.PreValue,
		OpId:       "qudongjie",
	}
	res, err := config.FccRpcClient.PrePublish(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res == nil || res.BaseRet == nil || res.BaseRet.Code != 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("res=%+v", res))
		return
	}
	c.JSON(http.StatusOK, "")
}

func Publish(c *gin.Context) {
	// swagger:operation POST /fcc/config/publish/ Fcc CommonReq
	//
	// publish config
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/PublishReq"
	// responses:
	//   "200":
	//     description: success
	//     schema:
	//       allOf:
	//       - "$ref": "#/definitions/BaseResponse"
	//       - type: object
	//         properties:
	//           data:
	//
	param := &models.PublishReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.PublishRequest{
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		ConfKey:    param.ConfKey,
		OpId:       "qudongjie",
	}
	res, err := config.FccRpcClient.Publish(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res == nil || res.BaseRet == nil || res.BaseRet.Code != 0 {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("res=%+v", res))
		return
	}
	c.JSON(http.StatusOK, "")
}

func GetConfig(c *gin.Context) {
	// swagger:operation GET /fcc/config/value/ Fcc GetConfigReq
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
	//             "$ref": "#/definitions/Config"
	//
	param := &models.GetConfigReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &fcc_serv.FetchConfigRequest{
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		ConfKey:    param.ConfKey,
	}
	res, err := config.FccRpcClient.FetchConfig(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res.Data)
}
