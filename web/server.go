package main

import (
	"buggmaker/common/storage"
	"buggmaker/web/router"
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {

	// init mongo
	err := storage.InitMongo()
	if err != nil {
		return
	}
	defer storage.MongoSession.Close()
	// init redis
	err = storage.InitRedis()
	if err != nil {
		return
	}
	defer storage.RedisS.Close()

	c := context.Background()
	result, err := storage.RedisS.Ping(c).Result()
	if err != nil {
		return
	}

	fmt.Println(result)

	err = storage.MongoSession.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	// init iris
	app := iris.New()
	// router
	router.Hub(app)
	// start
	app.Run(iris.Addr(":80"))
}
