package models

// swagger:parameters StarlingProjectListReq
type StarlingProjectListReq struct {
	// in:query
	Limit       int64  `json:"limit" form:"limit" binding:"required"`
	Offset      int64  `json:"offset" form:"offset"`
	ProjectKey  string `json:"project_key" form:"project_key"`
	ProjectName string `json:"project_name" form:"project_name"`
}

// swagger:parameters StarlingGroupListReq
type StarlingGroupListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key"`
	GroupName  string `json:"group_name" form:"group_name"`
}

// swagger:parameters StarlingOriginLgListReq
type StarlingOriginLgListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	LangKey    string `json:"conf_key" form:"conf_key"`
}

// swagger:parameters StarlingTransLgListReq
type StarlingTransLgListReq struct {
	// in:query
	Limit      int64  `json:"limit" form:"limit" binding:"required"`
	Offset     int64  `json:"offset" form:"offset"`
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	LangKey    string `json:"lang_key" form:"lang_key"`
}

// swagger:parameters StarlingGetConfigReq
type StarlingGetTransLgReq struct {
	// in:query
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	LangKey    string `json:"lang_key" form:"lang_key" binding:"required"`
	Lang       string `json:"lang" form:"lang" binding:"required"`
}

// swagger:model StarlingSaveProjectReq
type StarlingSaveProjectReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	ProjectName string `json:"project_name" form:"project_name"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model StarlingSaveGroupReq
type StarlingSaveGroupReq struct {
	ProjectKey  string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey    string `json:"group_key" form:"group_key" binding:"required"`
	GroupName   string `json:"group_name,omitempty" form:"group_name"`
	Description string `json:"description,omitempty" form:"description"`
	Status      int64  `json:"status,omitempty" form:"status"`
}

// swagger:model StarlingSaveOriginLgReq
type StarlingSaveOriginLgReq struct {
	ProjectKey string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey   string `json:"group_key" form:"group_key" binding:"required"`
	LangKey    string `json:"lang_key" form:"lang_key" binding:"required"`
	Lang       string `json:"lang" form:"lang" binding:"required"`
	OriginText string `json:"origin_text,omitempty" form:"origin_text"`
	Status     int64  `json:"status,omitempty" form:"status"`
}

// swagger:model StarlingSaveTransLgReq
type StarlingSaveTransLgReq struct {
	ProjectKey    string `json:"project_key" form:"project_key" binding:"required"`
	GroupKey      string `json:"group_key" form:"group_key" binding:"required"`
	LangKey       string `json:"lang_key" form:"lang_key" binding:"required"`
	Lang          string `json:"lang" form:"lang" binding:"required"`
	TranslateText string `json:"translate_text,omitempty" form:"translate_text"`
	Status        int64  `json:"status,omitempty" form:"status"`
}
