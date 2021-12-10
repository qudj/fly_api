package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qudj/fly_api/config"
	"github.com/qudj/fly_api/models"
	servbp "github.com/qudj/fly_lib/models/proto/fly_starling_serv"
	"net/http"
	"strings"
)

func StarlingProjectList(c *gin.Context) {
	// swagger:operation GET /starling/project/list/ Starling StarlingProjectListReq
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
	param := &models.StarlingProjectListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.FetchProjectsRequest{
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
	res, err := config.StarlingRpcClient.FetchProjects(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(0, res.Data)
}

func StarlingGroupList(c *gin.Context) {
	// swagger:operation GET /starling/group/list/ Starling StarlingGroupListReq
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
	param := &models.StarlingGroupListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.FetchGroupsRequest{
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
	res, err := config.StarlingRpcClient.FetchGroups(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func StarlingOriginLgList(c *gin.Context) {
	// swagger:operation GET /starling/origin/list/ Starling StarlingOriginLgListReq
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
	//             "$ref": "#/definitions/FetchOriginLgsRet"
	//
	param := &models.StarlingOriginLgListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.FetchOriginLgsRequest{
		Limit:      param.Limit,
		Offset:     param.Offset,
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		Filter:     map[string]string{},
	}
	if param.LangKey != "" {
		req.Filter["lang_key"] = param.LangKey
	}
	res, err := config.StarlingRpcClient.FetchOriginLgs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func StarlingTransLgList(c *gin.Context) {
	// swagger:operation GET /starling/translation/list/ Starling StarlingTransLgListReq
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
	//             "$ref": "#/definitions/FetchTransLgsRet"
	//
	param := &models.StarlingTransLgListReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.FetchTransLgsRequest{
		Limit:      param.Limit,
		Offset:     param.Offset,
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		LangKey:    param.LangKey,
		Filter:     map[string]string{},
	}
	res, err := config.StarlingRpcClient.FetchTransLgs(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"msg": "ok"})
		return
	}
	c.JSON(http.StatusOK, res.Data)
}

func StarlingSaveProject(c *gin.Context) {
	// swagger:operation POST /starling/project/save/ Starling CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/StarlingSaveProjectReq"
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
	param := &models.StarlingSaveProjectReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.SaveProjectRequest{
		Project: &servbp.Project{
			ProjectKey:  param.ProjectKey,
			ProjectName: param.ProjectName,
			Description: param.Description,
			Status:      param.Status,
			OpId:        "qudongjie",
			OpName:      "qudongjie",
		},
		SaveMode: servbp.SaveMode(param.SaveMode),
	}
	res, err := config.StarlingRpcClient.SaveProject(c, req)
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

func StarlingSaveGroup(c *gin.Context) {
	// swagger:operation POST /starling/group/save/ Starling CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/StarlingSaveGroupReq"
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
	param := &models.StarlingSaveGroupReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.SaveGroupRequest{
		Group: &servbp.Group{
			ProjectKey:  param.ProjectKey,
			GroupKey:    param.GroupKey,
			GroupName:   param.GroupName,
			Description: param.Description,
			Status:      param.Status,
			OpId:        "qudongjie",
			OpName:      "qudongjie",
		},
		SaveMode: servbp.SaveMode(param.SaveMode),
	}
	res, err := config.StarlingRpcClient.SaveGroup(c, req)
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

func StarlingSaveOriginLg(c *gin.Context) {
	// swagger:operation POST /starling/origin/save/ Starling CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/StarlingSaveOriginLgReq"
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
	param := &models.StarlingSaveOriginLgReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.SaveOriginLgRequest{
		OriginLang: &servbp.OriginLg{
			ProjectKey: param.ProjectKey,
			GroupKey:   param.GroupKey,
			LangKey:    param.LangKey,
			Lang:       param.Lang,
			OriginText: param.OriginText,
			Status:     param.Status,
			OpId:       "qudongjie",
			OpName:     "qudongjie",
		},
		SaveMode: servbp.SaveMode(param.SaveMode),
	}
	res, err := config.StarlingRpcClient.SaveOriginLg(c, req)
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

func StarlingSaveTransLg(c *gin.Context) {
	// swagger:operation POST /starling/translate/save/ Starling CommonReq
	//
	// Get group list
	//
	// ---
	// parameters:
	// - in: body
	//   name: body
	//   schema:
	//     "$ref": "#/definitions/StarlingSaveTransLgReq"
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
	param := &models.StarlingSaveTransLgReq{}
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.SaveTransLgRequest{
		TransLang: &servbp.TransLg{
			ProjectKey:    param.ProjectKey,
			GroupKey:      param.GroupKey,
			LangKey:       param.LangKey,
			Lang:          param.Lang,
			TranslateText: param.TranslateText,
			Status:        param.Status,
			OpId:          "qudongjie",
			OpName:        "qudongjie",
		},
		SaveMode: servbp.SaveMode(param.SaveMode),
	}
	res, err := config.StarlingRpcClient.SaveTransLg(c, req)
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

func StarlingGetTranslation(c *gin.Context) {
	// swagger:operation GET /starling/translation/ Starling StarlingGetTransLgReq
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
	//             "$ref": "#/definitions/TransLg"
	//
	param := &models.StarlingGetTransLgReq{}
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	req := &servbp.FetchTransLgsByKeyRequest{
		ProjectKey: param.ProjectKey,
		GroupKey:   param.GroupKey,
		LangKeys:   strings.Split(param.LangKey, ","),
		Lang:       param.Lang,
	}
	res, err := config.StarlingRpcClient.FetchTransLgsByKey(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, res.Data)
}
