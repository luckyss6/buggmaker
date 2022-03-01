package main

import (
	"buggmaker/storage/logs"
	"fmt"
)

// 日志参数
const (
	code    = "code"
	message = "message"
)

func main() {
	fmt.Println()
	// 初始化日志

	log, err := logs.NewLog()
	if err != nil {
		fmt.Println(err)
		return
	}

}
