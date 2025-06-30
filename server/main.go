package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hlinfocc/cySSHClient2/pkg/admin"
	"github.com/hlinfocc/cySSHClient2/pkg/config"
	"github.com/hlinfocc/cySSHClient2/pkg/crontab"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	"github.com/hlinfocc/cySSHClient2/pkg/version"
	"gopkg.in/yaml.v3"
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
	VersionFull    bool
	Web            bool
	Socket         bool
	Port           int
}

type Config struct {
	WebHook string `yaml:"webhook"` //钉钉机器人Webhook地址
}

/**
* 初始化命令行参数信息
 */
func initParams() Args {
	args := Args{}
	args.Port = 31918
	flag.BoolVar(&args.Initialization, "init", args.Initialization, "初始化数据信息")
	// flag.StringVar(&args.Profile, "c", args.Profile, "指定配置文件")
	flag.BoolVar(&args.Web, "w", args.Web, "启动web服务")
	flag.BoolVar(&args.Socket, "s", args.Web, "启动Socket服务")
	flag.IntVar(&args.Port, "p", args.Port, "指定web服务端口")
	flag.BoolVar(&args.Version, "v", args.Version, "显示版本信息")
	flag.BoolVar(&args.VersionFull, "V", args.VersionFull, "显示系统名称及版本信息")
	// 覆盖默认的Usage函数
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "[%s]用法: %s [选项]\n\n", version.SysTitle, os.Args[0])
		fmt.Fprintln(os.Stderr, "标准选项:")
		flag.PrintDefaults() // 打印默认帮助信息
		fmt.Fprintln(os.Stderr, "\n其他:")
		fmt.Fprintf(os.Stderr, "  * 直接运行也可以启动web服务: %s \n", os.Args[0])
	}
	flag.Parse()
	return args
}

/**
* 启动Socket服务
 */
func StartSocketServer() {
	socketPath := config.GetSocketPath()
	log.Printf("启动Socket服务……%s", socketPath)
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

func loadConfig() Config {
	var cfg Config
	isexist := utils.FileExists("/etc/cysshClient.yml")
	if isexist {
		// 读取 YAML 文件
		yamlFile, err := os.ReadFile("/etc/cysshClient.yml")
		if err != nil {
			log.Fatalf("Error reading YAML file: %v", err)
		}
		// 解析 YAML 文件
		err = yaml.Unmarshal(yamlFile, &cfg)
		if err != nil {
			log.Fatalf("Error unmarshalling YAML data: %v", err)
		}
		// log.Println(cfg)
	}
	return cfg
}

func main() {
	args := initParams()

	if args.Initialization {
		fmt.Println("初始化数据库")
		initdb.Init()
	} else if args.Version {
		// 显示版本信息
		fmt.Println(version.Simple())
	} else if args.VersionFull {
		// 显示系统名称及版本信息
		fmt.Println(version.Full())
	} else if args.Web {
		cfg := loadConfig()
		if cfg.WebHook != "" {
			go crontab.StartCrond(cfg.WebHook)
		}
		admin.StartWebServer(args.Port)
	} else if args.Socket {
		StartSocketServer()
	} else {
		fmt.Println("启动web服务……")
		cfg := loadConfig()
		if cfg.WebHook != "" {
			go crontab.StartCrond(cfg.WebHook)
		}
		// go StartSocketServer()
		admin.StartWebServer(args.Port)
	}
}
