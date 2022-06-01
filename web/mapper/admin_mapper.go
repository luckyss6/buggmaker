package mapper

import (
	"buggmaker/common/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AdminMapper struct {
	UserCollection mgo.Collection
}

//  FindList
//  @Description: 		查询用户列表
//  @receiver admin
//  @return userList	用户列表
//  @return err			错误

func (admin *AdminMapper) FindList() (userList []model.User, err error) {
	err = admin.UserCollection.Find(map[string]interface{}{}).All(&userList)
	return
}

//  CreateUser
//  @Description: 		创建用户
//  @receiver admin
//  @param username		用户名
//  @param password		密码
//  @return err			错误

func (admin *AdminMapper) CreateUser(username, password string) (err error) {

	err = admin.UserCollection.Insert(bson.M{
		"username": username,
		"password": password,
	})
	return err
}
