package main

import (
	"fmt"

	"github.com/hlinfocc/cySSHClient2/pkg/config"
)

func main() {
	fmt.Println("ok")
	dbpath := config.GetDbPath()
	fmt.Println("path:" + dbpath)
}
