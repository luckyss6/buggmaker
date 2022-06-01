package main

import (
	"buggmaker/common/log"
	"buggmaker/common/storage"
	"buggmaker/web/router"
	"github.com/kataras/iris/v12"
)

func main() {

	var (
		err error
	)

	if err = log.InitLog(); err != nil {
		return
	}

	// init mongo
	if err = storage.InitMongo(); err != nil {
		return
	}
	// init redis
	if err = storage.InitRedis(); err != nil {
		return
	}

	defer storage.RedisS.Close()

	// init iris
	app := iris.New()
	app.Logger()
	// router
	router.Hub(app)
	// start
	app.Run(iris.Addr(":80"))
}
