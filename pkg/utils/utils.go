package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"

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
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
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

func InputHostId() int {
	fmt.Print("请输入主机ID：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	hostId, err := strconv.Atoi(inputStr)
	errors.ThrowErrorMsg(err, "请输入正确的主机ID")
	return hostId
}
