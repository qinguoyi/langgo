package main

import (
	"StorageProxy/api"
	"StorageProxy/app/middleware"
	"StorageProxy/bootstrap"
	"StorageProxy/bootstrap/plugins"
)

func main() {
	// config log
	lgConfig := bootstrap.NewConfig()
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
