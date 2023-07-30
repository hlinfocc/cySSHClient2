package hostlist

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	_ "github.com/mattn/go-sqlite3"
)

type Hostlist struct {
	Id       int
	Host     string
	Username string
	Port     string
	Iskey    int
	Keypath  string
	Hostdesc string
}

func QueryHostlist() ([]*Hostlist, error) {
	initdb.CheckDBIsReadable()
	var hostLists []*Hostlist
	db := initdb.GetConn()

	sql := "select * from sshhostlist order by id"
	rows, err := db.Query(sql)
	defer db.Close()
	errors.CheckError(err)
	if err != nil {
		return hostLists, err
	}
	for rows.Next() {
		var id int
		var host string
		var username string
		var port string
		var iskey int
		var keypath string
		var hostdesc string
		err = rows.Scan(&id, &host, &username, &port, &iskey, &keypath, &hostdesc)
		errors.CheckError(err)
		var item *Hostlist
		item.Id = id
		item.Host = host
		item.Username = username
		item.Port = port
		item.Iskey = iskey
		item.Keypath = keypath
		item.Hostdesc = hostdesc
		hostLists = append(hostLists, item)
	}
	return hostLists, nil
}

func Insert(data *Hostlist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	stmt, err := db.Prepare("INSERT INTO sshhostlist(host,username, port, iskey, keypath,hostdesc) VALUES (?, ?, ?, ?, ?, ?)")
	defer db.Close()
	if err != nil {
		return errors.ReturnError(err)
	}
	_, err1 := stmt.Exec(data.Host, data.Username, data.Port, data.Iskey, data.Keypath, data.Hostdesc)
	if err1 != nil {
		return errors.ReturnError(err1)
	}
	return true
}
func Update(data *Hostlist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	stmt, err := db.Prepare("update sshhostlist set host=?,username=?, port=?, iskey=?, keypath=?,hostdesc=? where id=?")
	defer db.Close()
	if err != nil {
		return errors.ReturnError(err)
	}
	res, err1 := stmt.Exec(data.Host, data.Username, data.Port, data.Iskey, data.Keypath, data.Hostdesc, data.Id)
	if err1 != nil {
		return errors.ReturnError(err1)
	}
	affect, err2 := res.RowsAffected()
	if err2 != nil {
		return errors.ReturnError(err2)
	}
	if affect > 0 {
		return true
	} else {
		return false
	}
}
func Delete(id int) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	stmt, err := db.Prepare("delete from sshhostlist where id=?")
	defer db.Close()
	if err != nil {
		return errors.ReturnError(err)
	}
	res, err1 := stmt.Exec(id)
	if err1 != nil {
		return errors.ReturnError(err1)
	}
	affect, err2 := res.RowsAffected()
	if err2 != nil {
		return errors.ReturnError(err2)
	}
	if affect > 0 {
		return true
	} else {
		return false
	}
}
