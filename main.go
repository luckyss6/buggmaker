package main

import (
	_ "buggmaker/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"buggmaker/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
