package userinfolist

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

func QueryKeylist() ([]*entity.UserInfo, error) {
	initdb.CheckDBIsReadable()
	var userInfo []*entity.UserInfo
	db := initdb.GetConn()

	result := db.Order("id asc").Find(&userInfo)
	errors.CheckError(result.Error)
	return userInfo, nil
}
func QueryKeylistPage(page int, limit int) ([]*entity.UserInfo, int64, error) {
	initdb.CheckDBIsReadable()
	var userInfo []*entity.UserInfo
	db := initdb.GetConn()
	var total int64 = 0
	var resultError error
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query := db.Model(&entity.UserInfo{}).Select("*")
		query.Count(&total)
		result := query.Order("id asc").Limit(limit).Offset(offset).Scan(&userInfo)
		resultError = result.Error
	} else {
		result := db.Model(&entity.UserInfo{}).Select("*").Order("id asc").Scan(&userInfo)
		resultError = result.Error
	}
	return userInfo, total, resultError
}

func QueryOne(id int) (*entity.UserInfo, error) {
	initdb.CheckDBIsReadable()
	var userInfo *entity.UserInfo
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&userInfo)
	// errors.CheckError(result.Error)
	return userInfo, result.Error
}

func FetchByAccount(account string) (*entity.UserInfo, error) {
	initdb.CheckDBIsReadable()
	var userInfo *entity.UserInfo
	db := initdb.GetConn()
	result := db.Where("account = ?", account).First(&userInfo)
	return userInfo, result.Error
}

func Insert(data *entity.UserInfo) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Create(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	return true
}
func Update(data *entity.UserInfo) bool {
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
func Delete(id int) (bool, string) {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	userInfo, uerr := QueryOne(id)
	if uerr != nil {
		return errors.ReturnError(uerr), "删除失败"
	}
	if userInfo.Account == "admin" {
		return errors.ReturnError(uerr), "删除失败，管理员账号禁止删除"
	}
	result := db.Delete(&entity.UserInfo{}, id)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err), "删除失败"
	}
	affect := result.RowsAffected
	if affect > 0 {
		return true, "删除成功"
	} else {
		return false, "删除失败"
	}
}

func CountTotal() int64 {
	initdb.CheckDBIsReadable()
	db := initdb.GetConn()
	var total int64 = 0
	db.Model(&entity.UserInfo{}).Count(&total)
	return total
}
