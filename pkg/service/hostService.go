package service

import (
	"os"
	"strconv"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/dbhandle"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
)

func writeFile(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		errors.ThrowErrorMsg(err, "无法打开文件： "+filePath)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		errors.ThrowErrorMsg(err, "写入临时密钥失败")
	}
}

func HostSshHandle(hostId int) {
	hostInfo, err := dbhandle.QueryHostOneById(hostId)
	errors.CheckError(err)
	if hostInfo.Iskey == 1 {
		keysInfo, kerr := dbhandle.QueryKeyOneById(utils.String2Int(hostInfo.Keypath))
		errors.CheckError(kerr)
		homeDir, err := os.UserHomeDir()
		var keyDir string
		if err != nil {
			keyDir = "/tmp/.cyssh/"
		} else {
			keyDir = homeDir + "/.cyssh/"
		}
		chk, _ := utils.PathExists(keyDir)
		if !chk {
			os.Mkdir(keyDir, 0644)
		}
		keyFullPath := keyDir + "id_rsa" + strconv.Itoa(keysInfo.Id)
		writeFile(keyFullPath, keysInfo.Privatekey)
		utils.ExecuteCommandSSH("-i", keyFullPath, "-p", hostInfo.Port, hostInfo.Username+"@"+hostInfo.Host)
	} else {
		utils.ExecuteCommandSSH("-p", hostInfo.Port, hostInfo.Username+"@"+hostInfo.Host)
	}
}

func SyncKey2Host(hostId int, keyId int) {
	hostInfo, err := dbhandle.QueryHostOneById(hostId)
	errors.CheckError(err)
	keysInfo, kerr := dbhandle.QueryKeyOneById(utils.String2IntDefault(hostInfo.Keypath, keyId))
	errors.CheckError(kerr)
	homeDir, err := os.UserHomeDir()
	var keyDir string
	if err != nil {
		keyDir = "/tmp/.cyssh/"
	} else {
		keyDir = homeDir + "/.cyssh/"
	}
	chk, _ := utils.PathExists(keyDir)
	if !chk {
		os.Mkdir(keyDir, 0644)
	}
	keyFullPath := keyDir + "id_rsa" + strconv.Itoa(keysInfo.Id) + ".pub"
	writeFile(keyFullPath, keysInfo.Publickey)
	utils.ExecuteCommandAny("/usr/bin/ssh-copy-id", "-i", keyFullPath, "-p", hostInfo.Port, hostInfo.Username+"@"+hostInfo.Host)
}

func HostScpHandle(hostId int, local2Remote bool, source1 string, source2 string) {
	hostInfo, err := dbhandle.QueryHostOneById(hostId)
	errors.CheckError(err)
	if hostInfo.Iskey == 1 {
		keysInfo, kerr := dbhandle.QueryKeyOneById(utils.String2Int(hostInfo.Keypath))
		errors.CheckError(kerr)
		homeDir, err := os.UserHomeDir()
		var keyDir string
		if err != nil {
			keyDir = "/tmp/.cyssh/"
		} else {
			keyDir = homeDir + "/.cyssh/"
		}
		chk, _ := utils.PathExists(keyDir)
		if !chk {
			os.Mkdir(keyDir, 0644)
		}
		keyFullPath := keyDir + "id_rsa" + strconv.Itoa(keysInfo.Id)
		writeFile(keyFullPath, keysInfo.Privatekey)
		if local2Remote {
			utils.ExecuteCommandAny("/usr/bin/scp", "-i", keyFullPath, "-P", hostInfo.Port, source1, hostInfo.Username+"@"+hostInfo.Host+":"+source2)
		} else {
			utils.ExecuteCommandAny("/usr/bin/scp", "-i", keyFullPath, "-P", hostInfo.Port, hostInfo.Username+"@"+hostInfo.Host+":"+source1, source2)
		}
	} else {
		if local2Remote {
			utils.ExecuteCommandAny("/usr/bin/scp", "-P", hostInfo.Port, source1, hostInfo.Username+"@"+hostInfo.Host+":"+source2)
		} else {
			utils.ExecuteCommandAny("/usr/bin/scp", "-P", hostInfo.Port, hostInfo.Username+"@"+hostInfo.Host+":"+source1, source2)
		}
	}
}
