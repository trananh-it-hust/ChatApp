package response

import (
	"net/http"

	"github.com/trananh-it-hust/ChatApp/pkg/util"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginatedResponse struct {
	Item       interface{}     `json:"item"`
	Pagination util.Pagination `json:"pagination"`
}

type ListResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewPaginatedResponse(item interface{}, pagination util.Pagination) *PaginatedResponse {
	return &PaginatedResponse{
		Item:       item,
		Pagination: pagination,
	}
}

func NewListResponse(code int, message string, data interface{}) *ListResponse {
	return &ListResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse(data interface{}) *Response {
	return NewResponse(http.StatusOK, "success", data)
}

func NewPaginatedSuccessResponse(item interface{}, pagination util.Pagination) *PaginatedResponse {
	return NewPaginatedResponse(item, pagination)
}

func NewListSuccessResponse(data interface{}) *ListResponse {
	return NewListResponse(http.StatusOK, "success", data)
}
