package controller

import (
	"github.com/yb7/echoswg"
	"tnwl_download/service"
	"tnwl_download/util"
)

type ExcelController struct{}

func init() {
	controller := new(ExcelController)
	e := echoswg.NewApiGroup(util.EchoInstance, "excel", "/api/excel")
	e.GET("/test", controller.test, "测试")

}

func (*ExcelController) test() error {
	return service.ExcelServiceImpl.Test()
}