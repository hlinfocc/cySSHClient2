package utils

import (
	"bufio"
	"bytes"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/hlinfocc/cySSHClient2/pkg/datavo"
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

func RandStringBytes(n int, letter bool) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	if letter {
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

// 调用ssh-keygen生成ssh秘钥；
// 返回第一个为私钥，第二个为公钥
func Sshkeygen(pwd string) (string, string, error) {
	var tempPath string
	keyName := RandStringBytes(10, true)
	tempRandom, err := os.MkdirTemp("", "cyssh-*")
	if err == nil {
		tempPath = tempRandom
	} else {
		return "", "", errors.New("创建临时目录失败！")
	}
	cmdstr := fmt.Sprintf("/usr/bin/ssh-keygen -t rsa -b 4096 -f %s/%s -N \"%s\" -q", tempPath, keyName, pwd)
	fmt.Println("cmd:", cmdstr)
	c := cmd.NewCmd("/usr/bin/bash", "-c", cmdstr)
	<-c.Start()
	rs := c.Status().Stdout
	if len(rs) > 0 {
		return "", "", errors.New(strings.Join(rs, ";"))
	}
	var privatekey string
	var publickey string
	if FileExists(fmt.Sprintf("%s/%s", tempPath, keyName)) {
		privateContent, _ := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, keyName))
		privatekey = string(privateContent)
	}
	if FileExists(fmt.Sprintf("%s/%s", tempPath, keyName)) {
		publicContent, _ := os.ReadFile(fmt.Sprintf("%s/%s", tempPath, keyName))
		publickey = string(publicContent)
	}
	os.RemoveAll(tempPath)
	return privatekey, publickey, nil
}

func Sha3(data string) string {
	h := sha512.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	return string(hash)
}

func RemoteRequest(url string, params any, requestMethod string) bool {
	portPath := "/var/run/hlinfo-cyssh-server.port"
	if FileExists(portPath) {
		portData, perr := os.ReadFile(portPath)
		if perr != nil {
			return errors.ReturnError(perr)
		}
		reqUrl := fmt.Sprintf("http://127.0.0.1:%s/%s", string(portData), url)

		// 创建自定义客户端
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		var req *http.Request
		var err error

		// 根据请求方法创建不同的请求
		switch requestMethod {
		case "POST":
			jsonData, err := json.Marshal(params)
			if err != nil {
				return errors.ReturnError(err)
			}
			req, err = http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonData))
			if err != nil {
				return errors.ReturnError(err)
			}
			req.Header.Set("Content-Type", "application/json")

		case "GET", "DELETE":
			req, err = http.NewRequest(requestMethod, reqUrl, nil)
			if err != nil {
				return errors.ReturnError(err)
			}
			// 将参数添加到URL查询字符串中
			if params != nil {
				query := req.URL.Query()
				paramMap, ok := params.(map[string]string)
				if ok {
					for k, v := range paramMap {
						query.Add(k, v)
					}
					req.URL.RawQuery = query.Encode()
				}
			}

		default:
			return errors.ReturnError(errors.New("不支持的请求方法"))
		}

		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Request error:", err)
			return errors.ReturnError(err)
		}
		defer resp.Body.Close()

		// 处理响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read response error:", err)
			return errors.ReturnError(err)
		}
		var simpResp datavo.SimpResp
		err = json.Unmarshal(body, &simpResp)
		if err != nil {
			fmt.Println("解析JSON失败:", err)
			return errors.ReturnError(err)
		}
		if resp.StatusCode == 200 {
			if simpResp.Code == 200 {
				return true
			}
			return errors.ReturnError(errors.New(simpResp.Msg))
		}
		return errors.ReturnError(errors.New("请求失败"))
	}
	return errors.ReturnError(errors.New("后台服务未启动"))
}
