package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/onsi/gomega"
	"github.com/yb7/alilog"
	"github.com/yb7/echoswg"
	"tnwl_download/config"
	"tnwl_download/util"
)

func main() {

	gomega.RegisterFailHandler(func(message string, callerSkip ...int) {
		// do nothing
	})

	e := util.EchoInstance

	echoswg.ServeSwagger(e, echoswg.SwaggerConfig{
		UrlPrefix:   "/api",
		Title:       "Download Api",
		Description: "",
		CdnPrefix:   "https://img.cls.cn/statics/swagger-ui-4.10.3",
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} ${method} ${uri} ${status} cost:${latency_human} bytes:${bytes_in}->${bytes_out}}` + "\n",
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 9}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Use(util.EchoRecover)

	alilog.Infof("rest server started at port [%s]", config.C.Ports.Http)

	e.Logger.Fatal(e.Start(config.C.Ports.Http))

}
