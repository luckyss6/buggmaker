package router

import (
	"buggmaker/web/mapper"
	"fmt"
	"github.com/kataras/iris/v12"
)

func AdminRouter(party iris.Party) {
	var (
		admin = party.Party("/admin")
		aus   = adminService{adminDB: mapper.AdminMapper{}}
	)
	admin.Get("/", aus.Show)

}

type adminService struct {
	adminDB mapper.AdminMapper
}

func (admin *adminService) Show(ctx iris.Context) {
	_, err := admin.adminDB.FindList()
	if err != "" {
		fmt.Println(err)
		return
	}

}
