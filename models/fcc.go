package models

// swagger:parameters ProjectListReq
type ProjectListReq struct {
	// in:query
	Limit       int64  `json:"limit" form:"limit" binding:"required"`
	Offset      int64  `json:"offset" form:"offset"`
	ProjectKey  string `json:"project_key" form:"project_key"`
	ProjectName string `json:"project_name" form:"project_name"`
}

// swagger:parameters GroupListReq
type GroupListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key"`
	GroupName  string `json:"group_name" form:"group_name"`
}

// swagger:parameters ConfigListReq
type ConfigListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key"`
}

// swagger:parameters GetConfigReq
type GetConfigReq struct {
	// in:query
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
}

// swagger:model SaveProjectReq
type SaveProjectReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	ProjectName string `json:"project_name" form:"project_name"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model SaveGroupReq
type SaveGroupReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey    string `json:"group_key" form:"group_key" binding:"required"`
	GroupName   string `json:"group_name,omitempty" form:"group_name"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model SaveConfigReq
type SaveConfigReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey    string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey     string `json:"conf_key" form:"conf_key" binding:"required"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model PrePublishReq
type PrePublishReq struct {
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
	PreValue   string `json:"pre_value" form:"pre_value" binding:"required"`
}

// swagger:model PublishReq
type PublishReq struct {
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
}