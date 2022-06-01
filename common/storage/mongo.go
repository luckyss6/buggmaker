package storage

import (
	"fmt"
	mongo "gopkg.in/mgo.v2"
	"time"
)

var UserCollection *mongo.Collection

func InitMongo() (err error) {
	var (
		mongoConfig  *mongo.DialInfo
		MongoSession *mongo.Session
	)

	mongoConfig = &mongo.DialInfo{
		Addrs:    []string{"114.117.198.50"},
		Direct:   false,
		Timeout:  time.Second * 1,
		Username: "admin",
		Password: "123456",
	}

	if MongoSession, err = mongo.DialWithInfo(mongoConfig); err != nil {
		fmt.Printf("mgo dail error[%s]\n", err.Error())
		return
	}
	UserCollection = MongoSession.DB("user").C("user")
	return
}
