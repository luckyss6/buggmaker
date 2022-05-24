package main

import (
	"buggmaker/web/router"
	"github.com/kataras/iris/v12"
)

func main() {

	// init iris
	app := iris.New()

	// router
	router.Hub(app)

	// start
	app.Run(iris.Addr(":80"))
}
