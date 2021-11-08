package main

import (
	"fmt"

	"github.com/jianyuezhexue/MagicAdmin/magic"
)

func main() {
	fmt.Println("test")
	// fmt.Println(magic.Config)
	magic.Logger.Info("测试测试测试")
	magic.Logger.Debug("调试调试调试")
}
