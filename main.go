package main

import "github.com/kataras/iris/v12"

// 日志参数
// const (
// 	code    = "code"
// 	message = "message"
// )

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

func main() {
	// 启动项目
	app := iris.Default()

	// 加载日志
	app.Use(myMiddleware)

	app.Handle("Get", "/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Halo World"})
	})

	// 启动端口
	app.Run(iris.Addr(":80"))
}
