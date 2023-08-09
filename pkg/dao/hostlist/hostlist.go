package hostlist

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

func QueryHostlist() ([]*entity.HostlistAll, error) {
	initdb.CheckDBIsReadable()
	var result []*entity.HostlistAll
	db := initdb.GetConn()

	queryField := "sshhostlist.id, sshhostlist.host, sshhostlist.username,sshhostlist.port,sshhostlist.iskey,sshhostlist.keypath,sshhostlist.hostdesc, sshkeylist.keyname"
	res := db.Model(&entity.Sshhostlist{}).Select(queryField).Joins("left join sshkeylist on sshkeylist.id = sshhostlist.keypath").Scan(&result)
	errors.CheckError(res.Error)

	return result, nil
}

func QueryOne(id int) (*entity.Sshhostlist, error) {
	initdb.CheckDBIsReadable()
	var hostLists *entity.Sshhostlist
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&hostLists)
	errors.CheckError(result.Error)
	return hostLists, nil
}

func Insert(data *entity.Sshhostlist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Create(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	return true
}
func Update(data *entity.Sshhostlist) bool {
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

	result := db.Delete(&entity.Sshhostlist{}, id)
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
