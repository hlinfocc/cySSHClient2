package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const socketPath = "/tmp/cysshclient.sock"

type Resp struct {
	Code int
	Msg  string
	Data string
}

func main() {
	// service := ":1200"
	os.Remove(socketPath)
	tcpAddr, err := net.ResolveUnixAddr("unix", socketPath)
	checkError(err)
	listener, err := net.ListenUnix("unix", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go HandleServerConn(conn)
	}
}

func HandleServerConn(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()                                    // close connection before exit
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			fmt.Println(strings.TrimSpace(string(request[:read_len])))
			daytime := time.Now().String()
			rs := Resp{}
			rs.Code = 200
			rs.Msg = "获取成功:" + daytime
			rs.Data = "[{\"name\":\"1\"},{\"name\":\"2\"},{\"name\":\"3\"}]"
			v, _ := json.Marshal(rs)
			conn.Write([]byte(string(v)))
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
