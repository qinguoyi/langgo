package web

import (
	"github.com/gin-gonic/gin"
	"langgo/app/pkg/sqls"
)

/*****************************************************************/
/******** 从请求参数 或 url中获取参数，并整合成查询条件 *****************/
/*****************************************************************/

// ReqParams .
type ReqParams struct {
	c              *gin.Context
	sqls.Condition // 没有变量，仅仅是包含结构体，类似继承
}

// NewReqParams .
func NewReqParams(c *gin.Context) *ReqParams {
	return &ReqParams{
		c: c,
	}
}

func (q *ReqParams) getByPath(name string) string {
	return q.c.Param(name)
}

// EqByPath .
func (q *ReqParams) EqByPath(column string) *ReqParams {
	value := q.getByPath(column)
	if len(value) > 0 {
		q.Eq(column, value)
	}
	return q
}

// NotEqByPath .
func (q *ReqParams) NotEqByPath(column string) *ReqParams {
	value := q.getByPath(column)
	if len(value) > 0 {
		q.NotEq(column, value)
	}
	return q
}

func (q *ReqParams) getByQuery(name string) string {
	return q.c.Query(name)
}

// EqByQuery .
func (q *ReqParams) EqByQuery(column string) *ReqParams {
	value := q.getByQuery(column)
	if len(value) > 0 {
		q.Eq(column, value)
	}
	return q
}

// NotEqByQuery .
func (q *ReqParams) NotEqByQuery(column string) *ReqParams {
	value := q.getByQuery(column)
	if len(value) > 0 {
		q.NotEq(column, value)
	}
	return q
}

// Asc .
func (q *ReqParams) Asc(column string) *ReqParams {
	q.Orders = append(q.Orders, sqls.OrderByCol{Column: column, Asc: true})
	return q
}

// Desc .
func (q *ReqParams) Desc(column string) *ReqParams {
	q.Orders = append(q.Orders, sqls.OrderByCol{Column: column, Asc: false})
	return q
}

// PageByQuery .
func (q *ReqParams) PageByQuery() *ReqParams {
	if q.c == nil {
		return q
	}
	paging := GetPageInfo(q.c)
	q.Page(paging.Page, paging.Limit)
	return q
}

// Limit .
func (q *ReqParams) Limit(limit int) *ReqParams {
	q.Page(1, limit)
	return q
}

// Page .
func (q *ReqParams) Page(page, limit int) *ReqParams {
	if q.PageInfo == nil {
		q.PageInfo = &sqls.PageInfo{Page: page, Limit: limit}
	} else {
		q.PageInfo.Page = page
		q.PageInfo.Limit = limit
	}
	return q
}
