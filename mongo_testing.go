package main

import (
	"buggmaker/common/model"
	"context"
	"fmt"
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func main() {

	var (
		//userList []model.User
		checkUserList []model.User
		mongoSession  *mongo.Session
		err           error
	)

	context.TODO()

	mongoConfig := &mongo.DialInfo{
		Addrs:    []string{"114.117.198.50"},
		Direct:   false,
		Timeout:  time.Second * 1,
		Username: "admin",
		Password: "123456",
	}
	if mongoSession, err = mongo.DialWithInfo(mongoConfig); err != nil {
		return
	}

	userCollection := mongoSession.DB("user").C("user")
	//for i := 0; i < 100; i++ {
	//	u := model.User{
	//		Username: uuid.NewV4().String(),
	//		Password: uuid.NewV1().String(),
	//	}
	//
	//}
	//
	//for _, val := range userList {
	//	fmt.Println("userList:", val)
	//}

	//userList = append(userList, u)

	// 插入100条数据
	//if err = userCollection.Insert(bson.M{"username": uuid.NewV1().String(), "password": uuid.NewV4().String()}); err != nil {
	//	return
	//}
	fmt.Println(userCollection)
	if err = userCollection.Find(bson.M{}).All(&checkUserList); err != nil {
		return
	}
	for _, val := range checkUserList {
		fmt.Println(val)
	}
}
