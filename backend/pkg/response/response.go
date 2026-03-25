package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	CodeSuccess = 0

	CodeParamError   = 1001
	CodeMissingParam = 1002

	CodeUnauthorized    = 2001
	CodeTokenInvalid    = 2002
	CodeAccountDisabled = 2003

	CodeNotFound = 3001

	CodeInternalError      = 4001
	CodeServiceUnavailable = 4002
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

func ErrorWithStatus(c *gin.Context, httpStatus int, code int, message string) {
	c.JSON(httpStatus, Response{
		Code:    code,
		Message: message,
	})
}

func ParamError(c *gin.Context, message string) {
	Error(c, CodeParamError, message)
}

func MissingParam(c *gin.Context, message string) {
	Error(c, CodeMissingParam, message)
}

func Unauthorized(c *gin.Context, message string) {
	ErrorWithStatus(c, http.StatusUnauthorized, CodeUnauthorized, message)
}

func TokenInvalid(c *gin.Context, message string) {
	Error(c, CodeTokenInvalid, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, CodeNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	ErrorWithStatus(c, http.StatusInternalServerError, CodeInternalError, message)
}

func ServiceUnavailable(c *gin.Context, message string) {
	ErrorWithStatus(c, http.StatusServiceUnavailable, CodeServiceUnavailable, message)
}
