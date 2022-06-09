package router

import (
	"buggmaker/common/model"
	"buggmaker/web/mapper"
	"github.com/kataras/iris/v12"
)

func AdminRouter(party iris.Party) {
	var (
		admin = party.Party("/admin")
		aus   = adminService{adminDB: mapper.AdminMapper{}}
	)
	admin.Get("/", aus.ShowUserList)
	admin.Put("/{username:string}/{password:string}", aus.InsertUser)
}

type adminService struct {
	adminDB mapper.AdminMapper
}

type AdminRouteInterface interface {
	ShowUserList(ctx iris.Context)
	InsertUser(ctx iris.Context)
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

	ctx.JSON(result)
}

//  InsertUser
//  @Description: 		添加用户
//  @receiver admin
//  @param ctx

func (admin *adminService) InsertUser(ctx iris.Context) {
	var (
		err      error
		username = ctx.Params().GetStringDefault("username", "")
		password = ctx.Params().GetStringDefault("password", "")
	)
	if err = admin.adminDB.CreateUser(username, password); err != nil {
		return
	}
	ctx.JSON("ok")
}
