package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/hlinfocc/cySSHClient2/pkg/version"
)

/**
* 命令行参数结构体
 */
type Args struct {
	HostList   bool
	HostAdd    bool
	HostModify bool
	HostDel    bool
	KeyList    bool
	KeyAdd     bool
	KeyDel     bool
	KeyGen     bool
	KeySync    bool
	Version    bool
}

/**
* 初始化命令行参数信息
 */
func initParams() Args {
	args := Args{}
	flag.BoolVar(&args.HostList, "l", args.HostList, "列出所以主机")
	flag.BoolVar(&args.HostAdd, "i", args.HostAdd, "新增主机")
	flag.BoolVar(&args.HostModify, "m", args.HostModify, "编辑主机信息")
	flag.BoolVar(&args.HostDel, "d", args.HostDel, "删除主机")
	flag.BoolVar(&args.KeyList, "k", args.KeyList, "查询SSH证书列表")
	flag.BoolVar(&args.KeyAdd, "ki", args.KeyAdd, "添加一个本地SSH证书")
	flag.BoolVar(&args.KeyGen, "g", args.KeyGen, "调用ssh-keygen生成一个SSH RSA证书")
	flag.BoolVar(&args.KeyDel, "s", args.KeyDel, "删除一个SSH证书")
	flag.BoolVar(&args.KeySync, "r", args.KeySync, "同步证书公钥到远程主机")
	flag.BoolVar(&args.Version, "v", args.Version, "显示版本信息")

	flag.Parse()
	return args
}

func main() {
	// 命令行参数解析
	args := initParams()

	// fmt.Println(args)
	if args.HostList {
		fmt.Println("hostlist")
	} else if args.HostAdd {
		fmt.Println("adddddd")
	} else if args.Version {
		fmt.Println(version.Full())
	} else {

		if flag.NArg() > 0 {
			fmt.Println(flag.Args())
		}
	}
}

func reqBySocket() string {
	const socketPath = "/tmp/cysshclient.sock"
	tcpAddr, err := net.ResolveUnixAddr("unix", socketPath)
	checkError(err)
	conn, err := net.DialUnix("unix", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("timestamp2"))
	checkError(err)
	// result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result))
	// os.Exit(0)
	return string(result)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
