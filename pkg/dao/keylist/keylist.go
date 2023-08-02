package hostlist

import (
	"fmt"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/hostlist"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Sshkeylist struct {
	Id         int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Keyname    string
	Privatekey string
	Publickey  string
}

func QueryKeylist() ([]*Sshkeylist, error) {
	initdb.CheckDBIsReadable()
	var keylists []*Sshkeylist
	db := initdb.GetConn()

	result := db.Order("id asc").Find(&keylists)
	errors.CheckError(result.Error)
	return keylists, nil
}

func QueryOne(id string) (*Sshkeylist, error) {
	initdb.CheckDBIsReadable()
	var keyList *Sshkeylist
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&keyList)
	errors.CheckError(result.Error)
	return keyList, nil
}

func RenderKeyList() {
	list, err := QueryKeylist()
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

func Insert(data *Sshkeylist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Create(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	return true
}
func Update(data *Sshkeylist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	res := db.Save(&data)
	err := res.Error
	if err != nil {
		return errors.ReturnError(err)
	}

	affect := res.RowsAffected
	if affect > 0 {
		return true
	} else {
		return false
	}
}
func Delete(id int) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	var qty int64
	db.Model(&hostlist.Sshhostlist{}).Where("keypath = ?", id).Count(&qty)
	if qty > 0 {
		err := errors.New("该证书有主机已经绑定，禁止删除")
		return errors.ReturnError(err)
	}
	result := db.Delete(&Sshkeylist{}, id)
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
