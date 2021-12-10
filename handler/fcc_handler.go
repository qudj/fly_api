package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/config"
	"github.com/qudj/fly_api/models"
	"github.com/qudj/fly_lib/models/proto/fcc_serv"
	"net/http"
)

func FccProjectList(c *gin.Context) {
	// swagger:operation GET /fcc/project/list/ Fcc FccProjectListReq
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
	param := &models.FccProjectListReq{}
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

func FccGroupList(c *gin.Context) {
	// swagger:operation GET /fcc/group/list/ Fcc FccGroupListReq
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
	param := &models.FccGroupListReq{}
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

func FccConfigList(c *gin.Context) {
	// swagger:operation GET /fcc/config/list/ Fcc FccConfigListReq
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
	param := &models.FccConfigListReq{}
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

func FccSaveProject(c *gin.Context) {
	// swagger:operation POST /fcc/project/save/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/FccSaveProjectReq"
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
	param := &models.FccSaveProjectReq{}
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
			OpId:        "qudongjie",
			OpName:      "qudongjie",
		},
		SaveMode: fcc_serv.SaveMode(param.SaveMode),
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

func FccSaveGroup(c *gin.Context) {
	// swagger:operation POST /fcc/group/save/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/FccSaveGroupReq"
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
	param := &models.FccSaveGroupReq{}
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
			OpId:        "qudongjie",
			OpName:      "qudongjie",
		},
		SaveMode: fcc_serv.SaveMode(param.SaveMode),
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

func FccSaveConfig(c *gin.Context) {
	// swagger:operation POST /fcc/config/save/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/FccSaveConfigReq"
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
	param := &models.FccSaveConfigReq{}
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
			OpId:        "qudongjie",
			OpName:      "qudongjie",
		},
		SaveMode: fcc_serv.SaveMode(param.SaveMode),
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

func FccPrePublish(c *gin.Context) {
	// swagger:operation POST /fcc/config/pre_publish/ Fcc CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/FccPrePublishReq"
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
	param := &models.FccPrePublishReq{}
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

func FccPublish(c *gin.Context) {
	// swagger:operation POST /fcc/config/publish/ Fcc CommonReq
	//
	// publish config
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/FccPublishReq"
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
	param := &models.FccPublishReq{}
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

func FccGetConfig(c *gin.Context) {
	// swagger:operation GET /fcc/config/value/ Fcc FccGetConfigReq
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
	param := &models.FccGetConfigReq{}
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
