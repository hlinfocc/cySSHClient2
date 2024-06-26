package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hlinfocc/cySSHClient2/pkg/admin"
	"github.com/hlinfocc/cySSHClient2/pkg/config"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/version"
)

type Resp struct {
	Code int
	Msg  string
	Data string
}

/**
* 命令行参数结构体
 */
type Args struct {
	Initialization bool
	Profile        string
	Version        bool
	Web            bool
}

/**
* 初始化命令行参数信息
 */
func initParams() Args {
	args := Args{}
	flag.BoolVar(&args.Initialization, "init", args.Initialization, "初始化数据信息")
	flag.StringVar(&args.Profile, "c", args.Profile, "指定配置文件")
	flag.BoolVar(&args.Web, "w", args.Web, "启动web服务")
	flag.BoolVar(&args.Version, "v", args.Version, "显示版本信息")

	flag.Parse()
	return args
}

/**
* 启动Socket服务
 */
func StartServer() {
	socketPath := config.GetSocketPath()
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
	// 设置2分钟超时时间
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// 将最大请求长度设置为128B以防止DDos攻击
	request := make([]byte, 128)
	// 退出前关闭连接
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			// 客户端已关闭连接
			break
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

func main() {
	args := initParams()

	fmt.Println(args.Initialization)
	if args.Initialization {
		fmt.Println("初始化数据库")
		initdb.Init()
	} else if args.Version {
		fmt.Println(version.Full())
	} else if args.Web {
		admin.StartWebServer()
	} else {
		StartServer()
	}
}
