package db

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}