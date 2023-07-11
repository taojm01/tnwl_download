package util

import (
	"github.com/labstack/echo/v4"
	"github.com/yb7/alilog"
	"runtime"
	"tnwl_download/bizerrors"
)

var StackSize = 4 << 10 // 4 KB
func EchoRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				if bizError, ok := r.(*bizerrors.BizError); ok {
					c.JSON(bizError.HttpStatus, map[string]interface{}{
						"errno": bizError.Code(),
						"msg":   bizError.Error(),
					})
					return
				}
				err, ok := r.(error)
				if !ok {
					alilog.Error(err)
					//err = fmt.Errorf("%v", r)
				}
				runtime.StartTrace()
				stack := make([]byte, StackSize)
				length := runtime.Stack(stack, true)
				alilog.Errorf("[PANIC RECOVER] %v\n%s\n", err, string(stack[:length]))
				//if !config.DisablePrintStack {
				//  c.Logger().Printf("[PANIC RECOVER] %v %s\n", err, stack[:length])
				//}
				//c.Error(err)
			}
		}()
		return next(c)
	}
}
