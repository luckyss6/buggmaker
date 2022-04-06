package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// 日志参数
// const (
// 	code    = "code"
// 	message = "message"
// )

func main() {

	fmt.Println("halo")

	// 连接mongoDB
	mongoDB := InitMongoSession()

	client := mongoDB.DB("app").C("user")

	// 关闭连接
	defer mongoDB.Close()

	// 初始化gin
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.POST("/showUser", func(c *gin.Context) {
		c.JSON(200, client.Find(""))
	})
	// r.GET("/CreateUser", func(c *gin.Context) {
	// 	c.String(http.StatusOK,)
	// })

	r.Run("127.0.0.1:80")

}

func InitMongoSession() *mgo.Session {
	mHost := "114.117.198.50"
	mPort := "27017"
	//mDBName := "config"  //你要连接的表，也可以后面再选，都行
	mUsername := "admin"  //mongodb的账号
	mPassword := "123456" //mongodb的密码
	session, err := mgo.Dial(mHost + ":" + mPort)
	if err != nil {
		fmt.Println("mgo.Dial-error:", err)
		os.Exit(0)
	}
	session.SetMode(mgo.Eventual, true)
	myDB := session.DB("admin") //这里的关键是连接mongodb后，选择admin数据库，然后登录，确保账号密码无误之后，该连接就一直能用了
	//出现server returned error on SASL authentication step: Authentication failed. 这个错也是因为没有在admin数据库下登录
	err = myDB.Login(mUsername, mPassword)
	if err != nil {
		fmt.Println("Login-error:", err)
		os.Exit(0)
	}
	//myDB = session.DB(mDBName) //如果要在这里就选择数据库，这个myDB可以定义为全局变量
	return session
}
