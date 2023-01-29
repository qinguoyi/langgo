package sqls

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Condition struct {
	SelectCols []string
	Params     []ParamPair
	Orders     []OrderByCol
	PageInfo   *PageInfo
}

type ParamPair struct {
	Query string
	Args  []interface{}
}

type OrderByCol struct {
	Column string
	Asc    bool
}

// NewCondition 创建新条件
func NewCondition() *Condition {
	return &Condition{}
}

// Cols 查询字段
func (c *Condition) Cols(selectCols ...string) *Condition {
	if len(selectCols) > 0 {
		c.SelectCols = append(c.SelectCols, selectCols...)
	}
	return c
}

// Where 查询条件
func (c *Condition) Where(query string, args ...interface{}) *Condition {
	c.Params = append(c.Params, ParamPair{query, args})
	return c
}

// Eq 等值查询
func (c *Condition) Eq(column string, args ...interface{}) *Condition {
	c.Where(column+" = ?", args)
	return c
}

// NotEq 不等查询
func (c *Condition) NotEq(column string, args ...interface{}) *Condition {
	c.Where(column+" <> ?", args)
	return c
}

// Gt 大于查询
func (c *Condition) Gt(column string, args ...interface{}) *Condition {
	c.Where(column+" > ?", args)
	return c
}

// Gte 大于等于查询
func (c *Condition) Gte(column string, args ...interface{}) *Condition {
	c.Where(column+" >= ?", args)
	return c
}

// Lt 小于查询
func (c *Condition) Lt(column string, args ...interface{}) *Condition {
	c.Where(column+" < ?", args)
	return c
}

// Lte 小于等于查询
func (c *Condition) Lte(column string, args ...interface{}) *Condition {
	c.Where(column+" <= ?", args)
	return c
}

// Like 模糊查询
func (c *Condition) Like(column string, str string) *Condition {
	c.Where(column+" LIKE ?", "%"+str+"%")
	return c
}

// StartWith 开头匹配模糊查询
func (c *Condition) StartWith(column string, str string) *Condition {
	c.Where(column+" LIKE ?", str+"%")
	return c
}

// EndWith 结尾匹配模糊查询
func (c *Condition) EndWith(column string, str string) *Condition {
	c.Where(column+" LIKE ?", "%"+str)
	return c
}

// In 列表查询
func (c *Condition) In(column string, params interface{}) *Condition {
	c.Where(column+" in (?) ", params)
	return c
}

// NotIn 列表排除查询
func (c *Condition) NotIn(column string, params interface{}) *Condition {
	c.Where(column+" not in (?) ", params)
	return c
}

// Asc 升序
func (c *Condition) Asc(column string) *Condition {
	c.Orders = append(c.Orders, OrderByCol{Column: column, Asc: true})
	return c
}

// Desc 降序
func (c *Condition) Desc(column string) *Condition {
	c.Orders = append(c.Orders, OrderByCol{Column: column, Asc: false})
	return c
}

// Limit 查询限制
func (c *Condition) Limit(limit int) *Condition {
	c.Page(1, limit)
	return c
}

// Page 翻页
func (c *Condition) Page(page, limit int) *Condition {
	if c.PageInfo == nil {
		c.PageInfo = &PageInfo{Page: page, Limit: limit}
	} else {
		c.PageInfo.Page = page
		c.PageInfo.Limit = limit
	}
	return c
}

// Build 生成查询规则
func (c *Condition) Build(db *gorm.DB) *gorm.DB {
	ret := db

	// select
	if len(c.SelectCols) > 0 {
		ret = ret.Select(c.SelectCols)
	}

	// where
	if len(c.Params) > 0 {
		for _, param := range c.Params {
			ret = ret.Where(param.Query, param.Args...)
		}
	}

	// order
	if len(c.Orders) > 0 {
		for _, order := range c.Orders {
			if order.Asc {
				ret = ret.Order(order.Column + " ASC")
			} else {
				ret = ret.Order(order.Column + " DESC")
			}
		}
	}

	// limit
	if c.PageInfo != nil && c.PageInfo.Limit > 0 {
		ret = ret.Limit(c.PageInfo.Limit)
	}

	// offset
	if c.PageInfo != nil && c.PageInfo.Offset() > 0 {
		ret = ret.Offset(c.PageInfo.Offset())
	}

	return ret
}

// Find 查询数据
func (c *Condition) Find(db *gorm.DB, out interface{}) error {
	if err := c.Build(db).Find(out).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// FindOne 查询一条数据
func (c *Condition) FindOne(db *gorm.DB, out interface{}) error {
	if err := c.Limit(1).Build(db).First(out).Error; err != nil {
		return err
	}
	return nil
}

// Count 总数数据
func (c *Condition) Count(db *gorm.DB, model interface{}) int64 {
	ret := db.Model(model)

	// where
	if len(c.Params) > 0 {
		for _, query := range c.Params {
			ret = ret.Where(query.Query, query.Args...)
		}
	}

	var count int64
	if err := ret.Count(&count).Error; err != nil {
		logrus.Error(err)
	}
	return count
}
