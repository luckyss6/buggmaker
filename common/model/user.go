package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id_      bson.ObjectId `bson:"_id"`
	Password string        `json:"Password" bson:"password"`
	Username string        `json:"Username" bson:"username"`
}
