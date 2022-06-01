package main

import (
	"buggmaker/common/storage"
	"buggmaker/web/router"
	"github.com/kataras/iris/v12"
)

func main() {

	// init mongo
	err := storage.InitMongo()
	if err != nil {
		return
	}
	// init redis
	err = storage.InitRedis()
	if err != nil {
		return
	}
	defer storage.RedisS.Close()

	// init iris
	app := iris.New()
	// router
	router.Hub(app)
	// start
	app.Run(iris.Addr(":80"))
}
