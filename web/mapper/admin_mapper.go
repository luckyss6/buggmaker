package mapper

import (
	"buggmaker/common/model"
	"buggmaker/common/storage"
	"gopkg.in/mgo.v2/bson"
)

type AdminMapper struct {
}

type AdminInterface interface {
	FindList() (userList []model.User, err error)
	CreateUser(username, password string) error
	DelUser(username string) error
}

//  FindList
//  @Description: 		查询用户列表
//  @receiver admin
//  @return userList	用户列表
//  @return err			错误

func (admin *AdminMapper) FindList() (userList []model.User, err error) {
	err = storage.MongoSession.DB("user").C("user").Find(map[string]interface{}{}).All(&userList)
	return
}

//  CreateUser
//  @Description: 		创建用户
//  @receiver admin
//  @param username		用户名
//  @param password		密码
//  @return err			错误

func (admin *AdminMapper) CreateUser(username, password string) (err error) {

	err = storage.MongoSession.DB("user").C("user").Insert(bson.M{
		"username": username,
		"password": password,
	})
	return err
}

func (admin *AdminMapper) DelUser(username string) (err error) {
	return
}
