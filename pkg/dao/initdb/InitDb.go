package initdb

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/hlinfocc/cySSHClient2/pkg/config"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init() {
	dbpath := config.GetDbPath()
	fmt.Println("dbpath:" + dbpath)
	database, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	errors.CheckError(err)

	/* var tbHostList = `
	    CREATE TABLE IF NOT EXISTS "sshhostlist" (
	        "id"  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	        "host"  TEXT,
	        "username"  TEXT,
	        "port"  TEXT,
	        "iskey"  INTEGER,
	        "keypath"  TEXT,
	        "hostdesc"  TEXT
	    );
	` */
	err1 := database.AutoMigrate(&entity.Sshhostlist{})
	errors.WaringErr(err1)

	/* var tbsshkeylist = `
		CREATE TABLE IF NOT EXISTS sshkeylist (
			id  INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			keyname  TEXT,
			privatekey  TEXT,
	            publickey TEXT
	        );
			` */

	err2 := database.AutoMigrate(&entity.Sshkeylist{})
	errors.WaringErr(err2)

}

func GetConn() *gorm.DB {
	dbpath := config.GetDbPath()

	exists, err := utils.PathExists(dbpath)
	errors.CheckError(err)
	if !exists {
		errors.ThrowError("数据文件不存在，请初始化数据文件")
	}

	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	errors.CheckError(err)
	return db
}

func CheckDBIsWritable() {
	dbpath := config.GetDbPath()
	rs := utils.IsWritable(dbpath)
	if !rs {
		errors.ThrowError("数据文件【" + dbpath + "】对于当前用户没有可写权限")
	}
}

func CheckDBIsWritableBool() bool {
	dbpath := config.GetDbPath()
	rs := utils.IsWritable(dbpath)
	return rs
}

func CheckDBIsReadable() {
	dbpath := config.GetDbPath()
	rs := utils.IsReadable(dbpath)
	if !rs {
		errors.ThrowError("数据文件【" + dbpath + "】对于当前用户没有可读权限")
	}
}
