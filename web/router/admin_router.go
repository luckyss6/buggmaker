package router

import (
	"buggmaker/common/model"
	"buggmaker/web/mapper"
	"fmt"
	"github.com/kataras/iris/v12"
)

func AdminRouter(party iris.Party) {
	var (
		admin = party.Party("/admin")
		aus   = adminService{adminDB: mapper.AdminMapper{}}
	)
	admin.Get("/get", aus.ShowUserList)

}

type adminService struct {
	adminDB mapper.AdminMapper
}

//  ShowUserList
//  @Description:		展示用户列表
//  @receiver admin
//  @param ctx

func (admin *adminService) ShowUserList(ctx iris.Context) {
	var (
		err    error
		result []model.User
	)

	if result, err = admin.adminDB.FindList(); err != nil {
		return
	}

	fmt.Println("result:", result)

	ctx.JSON(result)
}

//  InsertUser
//  @Description: 		添加用户
//  @receiver admin
//  @param ctx

func (admin *adminService) InsertUser(ctx iris.Context) {

}
