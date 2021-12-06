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

// swagger:parameters ConfigListReq
type MiniConfigReq struct {
	// in:query
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	ConfKey    string `json:"conf_key" form:"conf_key" binding:"required"`
}