package main

import (
	"buggmaker/storage/cache"
	"fmt"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

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

	var redisConfig cache.Redis
	var redisClient cache.Res

	file, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return
	}

	// 从 yaml 中读取 redis 配置
	err = yaml.Unmarshal(file, &redisConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := redisClient.NewRedisClient(redisConfig)

	_, err = client.Ping().Result()
	if err != nil {
		return
	}

	app := iris.New()

	app.Use(myMiddleware)

	app.RegisterView(iris.HTML("./view", ".html"))

	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("hello.html")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	config := iris.WithConfiguration(iris.YAML("./iris.yaml"))
	app.Run(iris.Addr(":80"), config)
}
