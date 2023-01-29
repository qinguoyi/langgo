package web

import (
	"github.com/gin-gonic/gin"
	"langgo/app/pkg/sqls"
	"net/http"
)

/*
响应体，没有自定义业务错误代码，和httpCode保持统一
*/

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 自定义错误码
	Message string      `json:"message"` // 信息
	Data    interface{} `json:"data"`    // 数据
}

// Success 响应成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		"ok",
		data,
	})
}

// SuccessPageData 翻页数据
func SuccessPageData(c *gin.Context, results interface{}, page *sqls.PageInfo) {
	Success(c,
		&PageResult{
			Results: results,
			Page:    page,
		})
}

// SuccessCursorData 游标翻页数据
func SuccessCursorData(c *gin.Context, results interface{}, cursor string, hasMore bool) {
	Success(c,
		&CursorResult{
			Results: results,
			Cursor:  cursor,
			HasMore: hasMore,
		})
}

// ParamsError 参数错误
func ParamsError(c *gin.Context, msg string) {
	c.JSON(http.StatusUnprocessableEntity, Response{
		0,
		msg,
		"",
	})
}

// InternalError 内部错误
func InternalError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		0,
		msg,
		"",
	})
}

// UnAuthorization 未授权
func UnAuthorization(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Response{
		0,
		msg,
		"",
	})
}

// NotFoundResource 资源不存在
func NotFoundResource(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Response{
		0,
		msg,
		"",
	})
}
