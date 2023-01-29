package web

import (
	"github.com/gin-gonic/gin"
	"langgo/app/pkg/sqls"
	"strconv"
)

/*****************************************************************/
/**************** 从请求参数 或 url 或 请求体 中获取参数 ***************/
/*****************************************************************/

// GetByPath .
func GetByPath(c *gin.Context, name string) string {
	return c.Param(name)
}

// GetByQuery .
func GetByQuery(c *gin.Context, name string) string {
	return c.Query(name)
}

// GetIntDefaultByQuery .
func GetIntDefaultByQuery(c *gin.Context, name string, def int) (int, error) {
	if value, ok := c.GetQuery(name); ok {
		return strconv.Atoi(value)
	}
	return def, nil
}

// GetByJSONBody .
func GetByJSONBody(c *gin.Context, obj interface{}) error {
	return c.ShouldBindJSON(obj)
}

// GetPageInfo .
func GetPageInfo(c *gin.Context) *sqls.PageInfo {
	page, _ := GetIntDefaultByQuery(c, "page", 1)
	limit, _ := GetIntDefaultByQuery(c, "limit", 5)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 1000
	}
	return &sqls.PageInfo{Page: page, Limit: limit}
}
