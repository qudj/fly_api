package models

// swagger:parameters FccProjectListReq
type FccProjectListReq struct {
	// in:query
	Limit       int64  `json:"limit" form:"limit" binding:"required"`
	Offset      int64  `json:"offset" form:"offset"`
	ProjectKey  string `json:"project_key" form:"project_key"`
	ProjectName string `json:"project_name" form:"project_name"`
}

// swagger:parameters FccGroupListReq
type FccGroupListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key"`
	GroupName  string `json:"group_name" form:"group_name"`
}

// swagger:parameters FccConfigListReq
type FccConfigListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key"`
}

// swagger:parameters FccGetConfigReq
type FccGetConfigReq struct {
	// in:query
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
}

// swagger:model FccSaveProjectReq
type FccSaveProjectReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	ProjectName string `json:"project_name" form:"project_name"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model FccSaveGroupReq
type FccSaveGroupReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey    string `json:"group_key" form:"group_key" binding:"required"`
	GroupName   string `json:"group_name,omitempty" form:"group_name"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model FccSaveConfigReq
type FccSaveConfigReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey    string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey     string `json:"conf_key" form:"conf_key" binding:"required"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model FccPrePublishReq
type FccPrePublishReq struct {
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
	PreValue   string `json:"pre_value" form:"pre_value" binding:"required"`
}

// swagger:model FccPublishReq
type FccPublishReq struct {
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
}