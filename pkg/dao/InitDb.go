package dao

import (
	"database/sql"

	"github.com/hlinfocc/cySSHClient2/pkg/config"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	dbpath := config.GetDbPath()
	database, err := sql.Open("sqlite3", dbpath)
	errors.CheckError(err)

	var tbHostList = `
        CREATE TABLE IF NOT EXISTS "sshhostlist" (
            "id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
            "host"  TEXT,
            "username"  TEXT,
            "port"  TEXT,
            "iskey"  INTEGER,
            "keypath"  TEXT,
            "hostdesc"  TEXT
        )
    `
	stmt1, _ := database.Prepare(tbHostList)
	_, err1 := stmt1.Exec()
	errors.WaringErr(err1)

	var tbsshkeylist = `
        CREATE TABLE IF NOT EXISTS sshkeylist (
            id  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
            keyname  TEXT,
            privatekey  TEXT,
            publickey TEXT
        )
    `
	stmt2, _ := database.Prepare(tbsshkeylist)
	_, err2 := stmt2.Exec()
	errors.WaringErr(err2)

	database.Close()
}

func getConn() *sql.DB {
	dbpath := config.GetDbPath()

	exists, err := utils.PathExists(dbpath)
	errors.CheckError(err)
	if !exists {
		errors.ThrowError("数据文件不存在，请初始化数据文件")
	}

	db, err := sql.Open("sqlite3", dbpath)
	errors.CheckError(err)
	return db
}
