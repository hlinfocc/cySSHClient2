package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/dbhandle"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/service"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
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
func initParams() (Args, int) {
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
	hostId := -1
	if flag.NArg() > 0 {
		argsExt := flag.Args()
		for i := 0; i < len(argsExt); i++ {
			item := argsExt[i]
			itemInt := utils.String2Int(item)
			if itemInt > 0 {
				hostId = itemInt
				break
			}

		}
	}
	return args, hostId
}

func main() {
	// 命令行参数解析
	args, hostId := initParams()

	// fmt.Println(args)
	if args.HostList {
		dbhandle.RenderHostList()
	} else if args.HostAdd {
		dbhandle.AddHost()
	} else if args.HostModify {
		if hostId <= 0 {
			dbhandle.RenderHostList()
			hostId = utils.InputHostId()
		}
		dbhandle.UpdateHost(hostId)

	} else if args.HostDel {
		if hostId <= 0 {
			dbhandle.RenderHostList()
			hostId = utils.InputHostId()
		}
		dbhandle.DeleteHostById(hostId)
	} else if args.KeyList {
		dbhandle.RenderKeyList()
	} else if args.KeyAdd {
		dbhandle.AddKeyInfoLocal()
	} else if args.KeyGen {
		dbhandle.AddKeyInfo()
	} else if args.KeyDel {
		keyId := utils.InputInt("请输入ssh密钥对ID")
		dbhandle.DeleteKeyById(keyId, true)
	} else if args.KeySync {
		dbhandle.RenderHostList()
		hostId = utils.InputHostId()
		dbhandle.RenderKeyList()
		keyId := utils.InputInt("请输入ssh密钥对ID")
		service.SyncKey2Host(hostId, keyId)
	} else if args.Version {
		// 显示版本号
		fmt.Println(version.Full())
	} else {
		if flag.NArg() > 0 {
			if hostId <= 0 {
				dbhandle.RenderHostList()
				hostId = utils.InputHostId()
			}
		} else {
			dbhandle.RenderHostList()
			hostId = utils.InputHostId()
			fmt.Println(hostId)
		}
		service.HostSshHandle(hostId)
	}
}

func reqBySocket() string {
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
	// os.Exit(0)
	return string(result)
}
