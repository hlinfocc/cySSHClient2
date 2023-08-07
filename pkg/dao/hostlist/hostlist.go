package hostlist

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Sshhostlist struct {
	Id       int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Host     string
	Username string
	Port     string
	Iskey    int
	Keypath  string
	Hostdesc string
}
type HostlistAll struct {
	Id       int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Host     string
	Username string
	Port     string
	Iskey    int
	Keypath  string
	Hostdesc string
	Keyname  string
}

type Sshkeylist struct {
	Id         int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Keyname    string
	Privatekey string
	Publickey  string
}

func QueryHostlist() ([]*HostlistAll, error) {
	initdb.CheckDBIsReadable()
	var result []*HostlistAll
	db := initdb.GetConn()

	queryField := "sshhostlist.id, sshhostlist.host, sshhostlist.username,sshhostlist.port,sshhostlist.iskey,sshhostlist.keypath,sshhostlist.hostdesc, sshkeylist.keyname"
	res := db.Model(&Sshhostlist{}).Select(queryField).Joins("left join sshkeylist on sshkeylist.id = sshhostlist.keypath").Scan(&result)
	errors.CheckError(res.Error)

	return result, nil
}

func QueryOne(id string) (*Sshhostlist, error) {
	initdb.CheckDBIsReadable()
	var hostLists *Sshhostlist
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&hostLists)
	errors.CheckError(result.Error)
	return hostLists, nil
}

func RenderHostList() {
	hostlist, err := QueryHostlist()
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

func Insert(data *Sshhostlist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Create(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	return true
}
func Update(data *Sshhostlist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Save(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	affect := result.RowsAffected
	if affect > 0 {
		return true
	} else {
		return false
	}
}
func Delete(id int) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()

	result := db.Delete(&Sshhostlist{}, id)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	affect := result.RowsAffected
	if affect > 0 {
		return true
	} else {
		return false
	}
}

func QueryKeyLists() ([]*Sshkeylist, error) {
	initdb.CheckDBIsReadable()
	var keylists []*Sshkeylist
	db := initdb.GetConn()

	result := db.Order("id asc").Find(&keylists)
	errors.CheckError(result.Error)
	return keylists, nil
}

func RenderTbKeyList() {
	list, err := QueryKeyLists()
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

func QueryKeyOneById(id int) (*Sshkeylist, error) {
	initdb.CheckDBIsReadable()
	var keyList *Sshkeylist
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&keyList)
	errors.CheckError(result.Error)
	return keyList, nil
}

func AddHost() {
	host := utils.InputString("请输入主机名(域名或者IP)")
	username := utils.InputString("请输入用户名[默认:root]")
	hport := utils.InputPort("请输入端口[默认:22]")
	hostdesc := utils.InputString("请输入主机描述")
	iskeyok := utils.InputString("是否SSH密钥对登录[默认:No]?[Y/N]")

	hostdata := Sshhostlist{}

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
	if hport < 1 && hport > 65535 {
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
		// dd
		RenderTbKeyList()
		kid := utils.InputInt("请输入ssh密钥对ID")
		keyObj, keyObjErr := QueryKeyOneById(kid)
		if keyObjErr != nil {
			RenderTbKeyList()
			kid = utils.InputInt("请输入正确的ssh密钥对ID")
			keyObj, keyObjErr = QueryKeyOneById(kid)
			if keyObjErr != nil {
				fmt.Println("输入的ssh密钥对ID不正确，请重新操作")
				os.Exit(1)
			}
		}
		hostdata.Keypath = strconv.Itoa(keyObj.Id)
	}
	if initdb.CheckDBIsWritableBool() {
		Insert(&hostdata)
	} else {
		// 当前用户没有权限写入数据库采用socket
	}
}
