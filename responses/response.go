package responses

// swagger:model BaseResponse
type BaseResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
