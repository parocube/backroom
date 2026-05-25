package entity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/parocube/backroom/xuuid"
)

// Result represents a standard API response body.
type Result struct {
	HttpStatus   int         `json:"httpStatus"`
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	RedirectType string      `json:"redirectType,omitempty"`
	ShowType     string      `json:"showType,omitempty"`
	TraceID      uuid.UUID   `json:"traceId"`
	Data         interface{} `json:"data,omitempty"`
}

// page represents paginated response data with pagination metadata.
type page struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	PageSize int         `json:"pageSize"`
	Current  int         `json:"current"`
	Pages    int         `json:"pages"`
}

// CreateOK writes a successful creation response with the given data.
func CreateOK(c *gin.Context, data interface{}) {
	ok(c, http.StatusCreated, "create success", data)
}

// UpdateOK writes a successful update response with the given data.
func UpdateOK(c *gin.Context, data interface{}) {
	ok(c, http.StatusOK, "update success", data)
}

// QueryOK writes a successful query response with the given data.
func QueryOK(c *gin.Context, data interface{}) {
	ok(c, http.StatusOK, "query success", data)
}

// QueryPageOK writes a successful paginated query response with the given data.
// It includes pagination metadata in the response body.
func QueryPageOK(c *gin.Context, current int, pageSize int, total int64, data interface{}) {
	ok(c, http.StatusOK, "query success", page{
		Data:     data,
		Total:    total,
		PageSize: pageSize,
		Current:  current,
		Pages:    (int(total) + pageSize - 1) / pageSize,
	})
}

// DeleteOK writes a successful delete response.
func DeleteOK(c *gin.Context) {
	ok(c, http.StatusOK, "delete success", nil)
}

// Error builds an error response result with the given status and message.
// It includes showType, redirectType and traceId information.
func Error(c *gin.Context, httpStatus int, message string) {
	traceId, _ := xuuid.GetTraceIdFromHeader(c)
	c.AbortWithStatusJSON(httpStatus, Result{
		HttpStatus:   httpStatus,
		Success:      false,
		Message:      message,
		RedirectType: "",
		ShowType:     "",
		TraceID:      traceId,
		Data:         nil,
	})
}

// ok writes a successful JSON response with the given status, message and data.
func ok(c *gin.Context, httpStatus int, message string, data interface{}) {
	traceId, _ := xuuid.GetTraceIdFromHeader(c)
	c.JSON(httpStatus, Result{
		HttpStatus:   httpStatus,
		Success:      true,
		Message:      message,
		RedirectType: "",
		ShowType:     "",
		TraceID:      traceId,
		Data:         data,
	})
}
