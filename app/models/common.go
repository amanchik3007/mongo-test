package models

type BaseResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	MsgType string `json:"msgType,omitempty"`
}
