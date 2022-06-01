package mapper

import (
	"buggmaker/common/model"
	"buggmaker/common/storage"
)

type AdminMapper struct{}

//  FindList
//  @Description: 		查询用户列表
//  @receiver admin
//  @return userList	用户列表
//  @return err			错误

func (admin *AdminMapper) FindList() (userList []model.User, err error) {
	err = storage.UserCollection.Find(map[string]interface{}{}).All(&userList)
	return
}
