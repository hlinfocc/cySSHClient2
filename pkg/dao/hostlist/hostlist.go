package hostlist

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

func QueryHostlist(page int, limit int, description string, hostIp string, hostExtent bool) ([]*entity.HostlistAll, int64, error) {
	initdb.CheckDBIsReadable()
	var result []*entity.HostlistAll
	db := initdb.GetConn()
	var total int64 = 0
	queryField := "sshhostlist.id, sshhostlist.host, sshhostlist.username,sshhostlist.port,sshhostlist.iskey,sshhostlist.keypath,sshhostlist.hostdesc, sshkeylist.keyname"
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query := db.Model(&entity.Sshhostlist{}).Select(queryField).Joins("left join sshkeylist on sshkeylist.id = sshhostlist.keypath")
		// 添加description条件
		if description != "" {
			query = query.Where("sshhostlist.hostdesc LIKE ?", "%"+description+"%")
		}
		// 添加hostip条件
		if hostIp != "" {
			query = query.Where("sshhostlist.host LIKE ?", "%"+hostIp+"%")
		}

		// 根据hostExtent参数决定是否添加排除条件
		if hostExtent {
			query = query.Where("sshhostlist.id not in (select id from host_extent)")
		}

		query.Count(&total)
		res := query.Limit(limit).Offset(offset).Scan(&result)
		errors.CheckError(res.Error)
		// db.Model(&entity.Sshhostlist{}).Joins("left join sshkeylist on sshkeylist.id = sshhostlist.keypath").Count(&total)
	} else {
		res := db.Model(&entity.Sshhostlist{}).Select(queryField).Joins("left join sshkeylist on sshkeylist.id = sshhostlist.keypath").Scan(&result)
		errors.CheckError(res.Error)
	}

	return result, total, nil
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

func CountTotal() int64 {
	initdb.CheckDBIsReadable()
	db := initdb.GetConn()
	var total int64 = 0
	db.Model(&entity.Sshhostlist{}).Count(&total)
	return total
}
