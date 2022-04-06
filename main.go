package main

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	mongo "gopkg.in/mgo.v2"
)

// 日志参数
// const (
// 	code    = "code"
// 	message = "message"
// )

func main() {

	session, err := mongo.Dial("114.117.198.50")
	if err != nil {
		return
	}

	fmt.Println(session)

	fmt.Print("Start")

	r := gin.Default()
	r.Run()
	// // 初始化日志
	// router := gin.New()
	// err := router.Run("127.0.0.1:80")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// log, err := logs.NewLog()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
