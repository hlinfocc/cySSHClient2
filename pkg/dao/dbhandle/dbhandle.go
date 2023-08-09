package dbhandle

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/hostlist"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/keylist"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func RenderHostList() {
	hostlist, err := hostlist.QueryHostlist()
	errors.CheckError(err)
	t := table.NewWriter()
	header := table.Row{"ID", "Description", "Port", "Host", "ssh identity_file"}
	t.AppendHeader(header)
	var rows []table.Row
	for i := 0; i < len(hostlist); i++ {
		item := hostlist[i]
		rows = append(rows, table.Row{item.Id, item.Hostdesc, item.Port, fmt.Sprintf("%s@%s", item.Username, item.Host), item.Keyname})
	}
	t.AppendRows(rows)
	fmt.Println(t.Render())
}

func AddHost() (*entity.Sshhostlist, bool) {
	host := utils.InputString("请输入主机名(域名或者IP)")
	username := utils.InputString("请输入用户名[默认:root]")
	hport := utils.InputPort("请输入端口[默认:22]")
	hostdesc := utils.InputString("请输入主机描述")
	iskeyok := utils.InputString("是否SSH密钥对登录[默认:No]?[Y/N]")

	hostdata := entity.Sshhostlist{}

	if !utils.CheckNetAddr(host) {
		for {
			host = utils.InputString("请输入正确的主机名(域名或者IP)")
			if utils.CheckNetAddr(host) {
				break
			}
		}
	}
	hostdata.Host = host
	if len(username) <= 0 {
		username = "root"
	}
	hostdata.Username = username
	if hport < 1 || hport > 65535 {
		for {
			hport = utils.InputPort("请输入正确的端口[默认:22]")
			if hport >= 1 && hport <= 65535 {
				break
			}
		}
	}
	hostdata.Port = strconv.Itoa(hport)
	if len(hostdesc) <= 0 {
		for {
			hostdesc = utils.InputString("请输入主机描述")
			if len(hostdesc) > 0 {
				break
			}
		}
	}

	hostdata.Hostdesc = hostdesc

	if iskeyok == "Y" || iskeyok == "yes" || iskeyok == "y" {
		RenderKeyList()
		kid := utils.InputInt("请输入ssh密钥对ID")
		// keyObj, keyObjErr := QueryKeyOneById(kid)
		keyObj, keyObjErr := keylist.QueryOne(kid)
		if keyObjErr != nil {
			RenderKeyList()
			kid = utils.InputInt("请输入正确的ssh密钥对ID")
			keyObj, keyObjErr = keylist.QueryOne(kid)
			if keyObjErr != nil {
				fmt.Println("输入的ssh密钥对ID不正确，请重新操作")
				os.Exit(1)
			}
		}
		hostdata.Keypath = strconv.Itoa(keyObj.Id)
	}
	if initdb.CheckDBIsWritableBool() {
		rs := hostlist.Insert(&hostdata)
		return &hostdata, rs
	} else {
		return &hostdata, false
	}
}

func UpdateHost(id int) (*entity.Sshhostlist, bool) {
	data, ObjErr := hostlist.QueryOne(id)
	if ObjErr != nil {
		fmt.Println("输入的主机ID不正确，请重新操作")
		os.Exit(1)
	}
	host := utils.InputString("请输入新主机名(域名或者IP)[" + data.Host + "]")
	if utils.CheckNetAddr(host) {
		data.Host = host
	}

	username := utils.InputString("请输入用户名[" + data.Username + "]")
	if len(username) > 0 {
		data.Username = username
	}
	hport := utils.InputPort("请输入端口[" + data.Port + "]")
	if hport >= 1 && hport <= 65535 {
		data.Port = strconv.Itoa(hport)
	}
	hostdesc := utils.InputString("请输入主机描述[" + data.Hostdesc + "]")
	if len(hostdesc) > 0 {
		data.Hostdesc = hostdesc
	}
	var isKeyDefault string
	if len(data.Keypath) > 0 {
		isKeyDefault = "Y"
	} else {
		isKeyDefault = "N"
	}
	iskeyok := utils.InputString("是否SSH密钥对登录[" + isKeyDefault + "]?[Y/N]")
	if iskeyok == "Y" || iskeyok == "yes" || iskeyok == "y" {
		RenderKeyList()
		kid := utils.InputInt("请输入ssh密钥对ID")
		// keyObj, keyObjErr := QueryKeyOneById(kid)
		keyObj, keyObjErr := keylist.QueryOne(kid)
		if keyObjErr != nil {
			RenderKeyList()
			kid = utils.InputInt("请输入正确的ssh密钥对ID")
			keyObj, keyObjErr = keylist.QueryOne(kid)
			if keyObjErr != nil {
				fmt.Println("输入的ssh密钥对ID不正确，请重新操作")
				os.Exit(1)
			}
		}
		data.Keypath = strconv.Itoa(keyObj.Id)
	}
	if initdb.CheckDBIsWritableBool() {
		rs := hostlist.Update(data)
		return data, rs
	} else {
		return data, false
	}
}

func QueryHostOneById(id int) (*entity.Sshhostlist, error) {
	return hostlist.QueryOne(id)
}
func DeleteHostById(id int) bool {
	return hostlist.Delete(id)
}

func RenderKeyList() {
	list, err := keylist.QueryKeylist()
	errors.CheckError(err)
	t := table.NewWriter()
	header := table.Row{"ID", "ssh identity_file"}
	t.AppendHeader(header)
	var rows []table.Row
	for i := 0; i < len(list); i++ {
		item := list[i]
		rows = append(rows, table.Row{item.Id, item.Keyname})
	}
	t.AppendRows(rows)
	fmt.Println(t.Render())
}

func QueryKeyOneById(id int) (*entity.Sshkeylist, error) {
	return keylist.QueryOne(id)
}
func DeleteKeyById(id int) bool {
	return keylist.Delete(id)
}
