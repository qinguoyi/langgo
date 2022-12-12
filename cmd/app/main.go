package main

import (
	"langgo/api"
	"langgo/app/middleware"
	"langgo/bootstrap"
	"langgo/bootstrap/plugins"
)

func main() {
	// config log
	lgConfig := bootstrap.NewConfig("conf/config.yaml")
	lgLogger := bootstrap.NewLogger()

	// plugins DB Redis Minio
	plugins.NewPlugins()
	defer plugins.ClosePlugins()

	// middleware
	corsM := middleware.NewCors()
	traceL := middleware.NewTrace(lgLogger)
	requestL := middleware.NewRequestLog(lgLogger)

	// router
	engine := api.NewRouter(lgConfig, corsM, traceL, requestL)
	server := newHttpServer(lgConfig, engine)

	// app run-server
	app := newApp(lgConfig, lgLogger.Logger, server)
	app.runServer()
}
