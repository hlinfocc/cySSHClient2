package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/version"
)

/**
* 命令行参数结构体
 */
type Args struct {
	Local2Remote bool
	Remote2Local bool
	Source       string
	Target       string
	Version      bool
}

/**
* 初始化命令行参数信息
 */
func initParams() Args {
	args := Args{}
	flag.BoolVar(&args.Local2Remote, "l", args.Local2Remote, "从本地上传文件到远程服务器")
	flag.BoolVar(&args.Remote2Local, "r", args.Remote2Local, "从远程服务器复制文件到本地")
	flag.StringVar(&args.Source, "s", args.Source, "源文件或目录")
	flag.StringVar(&args.Target, "t", args.Target, "目标路径")
	flag.BoolVar(&args.Version, "v", args.Version, "显示版本信息")

	flag.Parse()
	return args
}

func main() {
	// 命令行参数解析
	args := initParams()

	fmt.Println(args)
	if args.Local2Remote && args.Remote2Local {
		errors.ThrowError("参数-l和-r不能同时存在")
	}
	if args.Local2Remote {
		fmt.Println("Local2Remote")
	} else if args.Remote2Local {
		fmt.Println("Remote2Local")
	} else if args.Version {
		fmt.Println(version.Full())
	} else {

		if flag.NArg() > 0 {
			fmt.Println(flag.Args())
		}
	}
}

func ReqBySocket() string {
	const socketPath = "/tmp/cysshclient.sock"
	tcpAddr, err := net.ResolveUnixAddr("unix", socketPath)
	errors.CheckError(err)
	conn, err := net.DialUnix("unix", nil, tcpAddr)
	errors.CheckError(err)
	_, err = conn.Write([]byte("timestamp2"))
	errors.CheckError(err)
	// result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	errors.CheckError(err)
	fmt.Println(string(result))
	return string(result)
}
