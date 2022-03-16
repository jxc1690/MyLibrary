package irisP

import (
	"github.com/kataras/iris/v12"
	. "github.com/kataras/iris/v12/context"
)

func NewVue(主路径 string) *iris.Application {
	app := iris.New()
	app.OnErrorCode(404, func(ctx Context) {
		ctx.StatusCode(200)
		if err := ctx.View("index.html"); err != nil {
			return
		}
	})
	app.HandleDir("/", 主路径)
	app.Get("/", func(context Context) {

	})
	return app
}
func Addr(addr string) iris.Runner {
	return iris.Addr(addr)
}
