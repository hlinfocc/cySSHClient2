package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/hlinfocc/cySSHClient2/pkg/config"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/table"
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
	t := table.NewWriter()
	header := table.Row{"ID", "IP", "Num", "PacketsRecv", "PacketLoss", "AvgRtt"}

	t.AppendHeader(header)
	for i := 1; i <= 5; i++ {
		row := table.Row{i, fmt.Sprintf("10.0.0.%v", i), i + 4, i, i, "AppendRow"}
		t.AppendRow(row)
	}
	fmt.Println(t.Render())
	// var hostId int
	input := utils.InputHostId()
	fmt.Println(input)
	// fmt.Println(hostId)
}
