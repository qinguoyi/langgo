package middleware

import (
	"StorageProxy/bootstrap"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
)

/*
打印请求和响应数据
*/

// RequestLog .
type RequestLog struct {
	Logger *bootstrap.LangGoLogger
}

// NewRequestLog .
func NewRequestLog(logger *bootstrap.LangGoLogger) *RequestLog {
	return &RequestLog{
		Logger: logger,
	}
}

// CustomResponseWriter _
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write _
func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString _
func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Handler request response日志打印 接管gin的默认日志
func (r RequestLog) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 打印请求 日志打印推荐用zap的方法替换fmt.Sprintf
		req := ""
		switch strings.ToUpper(c.Request.Method) {
		case "GET":
			req = fmt.Sprintf("RequestInfo: method=%s, url=%s, user_id=%s, form=%s",
				c.Request.Method, c.Request.URL, c.Request.Header["user-id"], c.Request.Form)
		case "POST", "PATCH", "DELETE", "PUT":
			req = fmt.Sprintf("RequestInfo: method=%s, url=%s, user_id=%s, form=%s, post_form=%s, body=%s",
				c.Request.Method, c.Request.URL, c.Request.Header["user-id"], c.Request.Form, c.Request.PostForm,
				c.Request.Body)
		default:
			req = fmt.Sprintf("RequestInfo: method=%s, url=%s, user_id=%s, form=%s, post_form=%s, body=%s",
				c.Request.Method, c.Request.URL, c.Request.Header["user-id"], c.Request.Form, c.Request.PostForm,
				c.Request.Body)
		}
		r.Logger.WithContext(c).Info(req)

		blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		// 处理业务
		c.Next()
		cost := time.Since(start)

		// 打印响应
		r.Logger.WithContext(c).Info("RequestInfo",
			zap.String("ip", c.ClientIP()),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("cost", cost),
			zap.String("data", blw.body.String()),
		)
	}
}
