package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/hlinfocc/cySSHClient2/pkg/config"
)

func main() {
	fmt.Println("ok")
	dbpath := config.GetDbPath()
	fmt.Println("path:" + dbpath)
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("ok")
	}
	username := currentUser.Username

	fmt.Printf("Username is: %s\n", username)
}
