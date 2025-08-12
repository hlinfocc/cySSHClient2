package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

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
	Local2Remote bool
	Remote2Local bool
	Source       string
	Target       string
	Version      bool
	VersionFull  bool
}

/**
* 初始化命令行参数信息
 */
func initParams() Args {
	args := Args{}
	flag.BoolVar(&args.Local2Remote, "l", args.Local2Remote, "从本地上传文件到远程服务器")
	flag.BoolVar(&args.Remote2Local, "r", args.Remote2Local, "从远程服务器复制文件到本地")
	flag.BoolVar(&args.Version, "v", args.Version, "显示版本信息")
	flag.BoolVar(&args.VersionFull, "V", args.VersionFull, "显示系统名称及版本信息")
	// 覆盖默认的Usage函数
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "用法: %s [选项]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "标准选项:")
		flag.PrintDefaults() // 打印默认帮助信息
		fmt.Fprintln(os.Stderr, "\n示例信息:")
		fmt.Fprintf(os.Stderr, "  * 示例1本地传到远程主机（需要选择主机）: %s -l ./example.txt /opt/\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  * 示例2本地传到远程主机[1:是直接指定主机ID]: %s -l ./example.txt 1:/opt/\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  * 示例3远程主机传到本地（需要选择主机）: %s -r /opt/ ./example.txt \n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  * 示例3远程主机传到本地[1:是直接指定主机ID]: %s -r 1:/opt/example.txt ./ \n", os.Args[0])
	}
	flag.Parse()
	return args
}

func main() {
	// 命令行参数解析
	args := initParams()

	// fmt.Printf("args:%v\n", args)
	if args.Local2Remote && args.Remote2Local {
		errors.ThrowError("参数-l和-r不能同时存在")
	}
	if len(flag.Args()) != 2 {
		errors.ThrowError(fmt.Sprintf("参数不正确，使用：%s -? 查看选项信息", os.Args[0]))
	}
	if args.Local2Remote {
		re := regexp.MustCompile(`^(\d+):(.*)$`)
		// 查找匹配项
		match := re.FindStringSubmatch(flag.Args()[1])
		var source1 = flag.Args()[0]
		var source2 = flag.Args()[1]
		var hostId = -1
		if len(match) > 1 {
			hostId = utils.String2Int(match[1])
			source2 = match[2]
		} else {
			dbhandle.RenderHostList()
			hostId = utils.InputHostId()
		}
		service.HostScpHandle(hostId, true, source1, source2)
	} else if args.Remote2Local {
		re := regexp.MustCompile(`^(\d+):(.*)$`)
		// 查找匹配项
		match := re.FindStringSubmatch(flag.Args()[0])
		var source1 = flag.Args()[0]
		var source2 = flag.Args()[1]
		var hostId = -1
		if len(match) > 1 {
			hostId = utils.String2Int(match[1])
			source1 = match[2]
		} else {
			dbhandle.RenderHostList()
			hostId = utils.InputHostId()
		}
		service.HostScpHandle(hostId, false, source1, source2)
	} else if args.Version {
		fmt.Println(version.Simple())
	} else if args.VersionFull {
		fmt.Println(version.Full())
	} else {
		errors.ThrowError(fmt.Sprintf("参数不正确，使用：%s -? 查看选项信息", os.Args[0]))
	}
}
