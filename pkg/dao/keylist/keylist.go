package hostlist

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	_ "github.com/mattn/go-sqlite3"
)

type Keylist struct {
	Id         int
	Keyname    string
	Privatekey string
	Publickey  string
}

func QueryKeylist() ([]*Keylist, error) {
	initdb.CheckDBIsReadable()
	var keylists []*Keylist
	db := initdb.GetConn()

	sql := "select * from sshkeylist order by id"
	rows, err := db.Query(sql)
	defer db.Close()
	errors.CheckError(err)
	if err != nil {
		return keylists, err
	}
	for rows.Next() {
		var id int
		var keyname string
		var privatekey string
		var publickey string
		err = rows.Scan(&id, &keyname, &privatekey, &publickey)
		errors.CheckError(err)
		var item *Keylist
		item.Id = id
		item.Keyname = keyname
		item.Privatekey = privatekey
		item.Publickey = publickey
		keylists = append(keylists, item)
	}
	return keylists, nil
}

func Insert(data *Keylist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	stmt, err := db.Prepare("INSERT INTO sshkeylist(keyname,privatekey, publickey) VALUES (?, ?, ?)")
	defer db.Close()
	if err != nil {
		return errors.ReturnError(err)
	}
	_, err1 := stmt.Exec(data.Keyname, data.Privatekey, data.Publickey)
	if err1 != nil {
		return errors.ReturnError(err1)
	}
	return true
}
func Update(data *Keylist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	stmt, err := db.Prepare("update sshkeylist set keyname=?,privatekey=?, publickey=? where id=?")
	defer db.Close()
	if err != nil {
		return errors.ReturnError(err)
	}
	res, err1 := stmt.Exec(data.Keyname, data.Privatekey, data.Publickey, data.Id)
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
	stmt, err := db.Prepare("delete from sshkeylist where id=?")
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
