package mapper

import "buggmaker/common/model"

type AdminMapper struct {
}

func (admin *AdminMapper) FindList() (userList []model.User, err string) {
	err = ccccDB.C("cccc").Find(nil).All(&userList).Error()
	return
}
