package hostextent

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

func Querylist(page int, limit int, izCrond int) ([]*entity.HostExtent, int64, error) {
	initdb.CheckDBIsReadable()
	var result []*entity.HostExtent
	db := initdb.GetConn()
	var total int64 = 0
	if page > 0 || limit > 0 {
		offset := (page - 1) * limit
		query := db.Model(&entity.HostExtent{})
		// 添加izCrond条件
		if izCrond == 0 || izCrond == 1 {
			query = query.Where("izCrond = ?", izCrond)
		}
		query.Count(&total)
		res := query.Limit(limit).Offset(offset).Scan(&result)
		errors.CheckError(res.Error)
	} else {
		res := db.Model(&entity.HostExtent{}).Scan(&result)
		errors.CheckError(res.Error)
	}

	return result, total, nil
}

func QueryOne(id int) (*entity.HostExtent, error) {
	initdb.CheckDBIsReadable()
	var hostLists *entity.HostExtent
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&hostLists)
	errors.CheckError(result.Error)
	return hostLists, nil
}

func Insert(data *entity.HostExtent) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Create(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	return true
}
func Update(data *entity.HostExtent) bool {
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

	result := db.Delete(&entity.HostExtent{}, id)
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
	db.Model(&entity.HostExtent{}).Count(&total)
	return total
}
