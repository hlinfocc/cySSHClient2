package keylist

import (
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
)

func QueryKeylist() ([]*entity.Sshkeylist, error) {
	initdb.CheckDBIsReadable()
	var keylists []*entity.Sshkeylist
	db := initdb.GetConn()

	result := db.Order("id asc").Find(&keylists)
	errors.CheckError(result.Error)
	return keylists, nil
}
func QueryKeylistPage(page int, limit int) ([]*entity.Sshkeylist, int64, error) {
	initdb.CheckDBIsReadable()
	var keylists []*entity.Sshkeylist
	db := initdb.GetConn()
	var total int64 = 0
	var resultError error
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query := db.Model(&entity.Sshkeylist{}).Select("id,keyname")
		query.Count(&total)
		result := query.Order("id asc").Limit(limit).Offset(offset).Scan(&keylists)
		resultError = result.Error
	} else {
		result := db.Model(&entity.Sshkeylist{}).Select("id,keyname").Order("id asc").Scan(&keylists)
		resultError = result.Error
	}
	return keylists, total, resultError
}

func QueryOne(id int) (*entity.Sshkeylist, error) {
	initdb.CheckDBIsReadable()
	var keyList *entity.Sshkeylist
	db := initdb.GetConn()
	result := db.Where("id = ?", id).First(&keyList)
	// errors.CheckError(result.Error)
	return keyList, result.Error
}

func Insert(data *entity.Sshkeylist) bool {
	initdb.CheckDBIsWritable()
	db := initdb.GetConn()
	result := db.Create(&data)
	err := result.Error
	if err != nil {
		return errors.ReturnError(err)
	}
	return true
}
func Update(data *entity.Sshkeylist) bool {
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
	var qty int64
	db.Model(entity.Sshhostlist{}).Where("keypath = ?", id).Count(&qty)
	if qty > 0 {
		err := errors.New("该证书有主机已经绑定，禁止删除")
		return errors.ReturnError(err), "该证书有主机已经绑定，禁止删除"
	}
	result := db.Delete(&entity.Sshkeylist{}, id)
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
	db.Model(&entity.Sshkeylist{}).Count(&total)
	return total
}
