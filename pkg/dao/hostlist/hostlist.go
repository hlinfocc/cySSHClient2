package hostlist

import (
	"fmt"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/mattn/go-sqlite3"
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
