package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"syscall"

	"github.com/go-cmd/cmd"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但不是不存在的错误，所以把这个错误原封不动的返回
}

// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false
	}
	//文件存在，判断是否可读
	return IsReadable(path)
}

// 判断所给文件是否有可执行权限
func IsExecute(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	fileMode := fileInfo.Mode()
	perm := fileMode.Perm()
	flag := perm & os.FileMode(73)
	if uint32(flag) == uint32(73) {
		return true, nil
	} else {
		return false, err
	}
}

func IsWritable(f string) bool {
	err := syscall.Access(f, syscall.O_RDWR)
	if err != nil {
		return false
	} else {
		return true
	}
}
func IsReadable(f string) bool {
	err := syscall.Access(f, syscall.O_RDONLY)
	if err != nil {
		return false
	} else {
		return true
	}
}

func InputHostId() int {
	fmt.Print("请输入主机ID：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	hostId, err := strconv.Atoi(inputStr)
	errors.ThrowErrorMsg(err, "请输入正确的主机ID")
	return hostId
}
func InputInt(msg string) int {
	if len(msg) == 0 {
		msg = "请输入"
	}
	fmt.Printf("%s：", msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	res, err := strconv.Atoi(inputStr)
	errors.ThrowErrorMsg(err, "只能输入数字^_^")
	return res
}

func InputPort(msg string) int {
	if len(msg) == 0 {
		msg = "请输入端口"
	}
	fmt.Printf("%s：", msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	if len(inputStr) == 0 {
		inputStr = "22"
	}
	res, err := strconv.Atoi(inputStr)
	errors.ThrowErrorMsg(err, "只能输入数字^_^")
	return res
}

func InputString(msg string) string {
	if len(msg) == 0 {
		msg = "请输入"
	}
	fmt.Printf("%s：", msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	return inputStr
}

func ExecuteCommandCheckLocale(cmdstr string) int {
	c := cmd.NewCmd("/usr/bin/bash", "-c", cmdstr)
	<-c.Start()
	rs := c.Status().Stdout[0]
	qty, err := strconv.Atoi(rs)
	if err != nil {
		return -1
	} else {
		return qty
	}
}

func IsSupportZhCn() {
	var localeLangArr = [4]string{"en_US.UTF-8 UTF-8", "zh_CN.UTF-8 UTF-8", "zh_CN.GBK GBK", "zh_CN GB2312"}
	langQty := -1
	for i := 0; i < len(localeLangArr); i++ {
		cmdstr := fmt.Sprintf("locale -a | grep \"%s\" | wc -l", localeLangArr[i])
		qty := ExecuteCommandCheckLocale(cmdstr)
		if qty > 0 {
			langQty++
		}
	}
}

func ValiedIP(ip string) bool {
	regex := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)
	return regex.MatchString(ip)
}

func CheckNetAddr(addr string) bool {
	if len(addr) <= 0 {
		return false
	}
	if ValiedIP(addr) {
		ipnet, err := net.LookupIP(addr)
		if err != nil {
			return false
		}
		if len(ipnet) == 0 {
			return false
		} else {
			return true
		}
	} else {
		ipnet, err := net.LookupHost(addr)
		if err != nil {
			return false
		}
		if len(ipnet) == 0 {
			return false
		} else {
			return true
		}
	}
}

func String2Int(s string) int {
	if len(s) == 0 {
		return -1
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return res
}
