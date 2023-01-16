package api

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	v0 "langgo/api/v0"
	"langgo/app/middleware"
	"langgo/config"
	"langgo/docs"
	_ "langgo/docs"
)

// NewRouter .
func NewRouter(
	conf *config.Configuration,
	corsM *middleware.Cors,
	traceL *middleware.TraceLog,
	requestL *middleware.RequestLog,
) *gin.Engine {
	if conf.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	// 跨域 trace-id 日志
	router.Use(corsM.Handler(), traceL.Handler(), requestL.Handler())

	// 静态资源
	router.StaticFile("/assets", "../../static/image/back.png")

	// 动态资源 注册 api 分组路由
	setApiGroupRoutes(router)

	// swag docs
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return router
}

func setApiGroupRoutes(
	router *gin.Engine,
) *gin.RouterGroup {
	group := router.Group("/api/storage")
	routerV0 := group.Group("/v0")
	{
		routerV0.GET("/ping", v0.Test)
	}
	return group
}
