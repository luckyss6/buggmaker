package router

import (
	"github.com/kataras/iris/v12"
)

func AdminRouter(party iris.Party) {
	var (
		admin = party.Party("/admin")
	)
	admin.Get("/")

}

func Show() {

}
