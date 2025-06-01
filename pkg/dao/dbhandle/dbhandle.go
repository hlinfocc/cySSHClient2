package dbhandle

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	hostextent "github.com/hlinfocc/cySSHClient2/pkg/dao/hostExtent"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/hostlist"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/initdb"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/keylist"
	userinfolist "github.com/hlinfocc/cySSHClient2/pkg/dao/userInfoList"
	"github.com/hlinfocc/cySSHClient2/pkg/datavo"
	"github.com/hlinfocc/cySSHClient2/pkg/errors"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	sm3utils "github.com/hlinfocc/cySSHClient2/pkg/utils/sm3Utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func RenderHostList() {
	hostlist, _, err := hostlist.QueryHostlist(0, 0, "", "", false)
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

func QueryHostsList(page int, limit int, desc string, hostIp string, hostExtent bool) datavo.HostResp {
	hostlist, total, err := hostlist.QueryHostlist(page, limit, desc, hostIp, hostExtent)
	errors.CheckError(err)
	res := datavo.HostResp{
		Code:  200,
		Msg:   "获取成功",
		Data:  hostlist,
		Count: total,
	}
	return res
}

func Save(hostdata *entity.Sshhostlist, izInsert bool) datavo.HostResp {
	var rs bool
	res := datavo.HostResp{
		Code:  200,
		Msg:   "保存成功",
		Count: 1,
	}
	if izInsert {
		rs = hostlist.Insert(hostdata)
	} else {
		if hostdata.Id == 0 {
			res.Code = 500
			res.Msg = "保存失败，参数异常，ID不能为空"
			res.Count = 0
			return res
		}
		rs = hostlist.Update(hostdata)
	}

	if !rs {
		res.Code = 500
		res.Msg = "保存失败"
		res.Count = 0
	}
	return res
}

func AddHost() (*entity.Sshhostlist, bool) {
	host := utils.InputString("请输入主机名(域名或者IP)")
	username := utils.InputString("请输入用户名[默认:root]")
	hport := utils.InputPort("请输入端口[默认:22]")
	hostdesc := utils.InputString("请输入主机描述")
	iskeyok := utils.InputString("是否SSH密钥对登录[默认:No]?[Y/N]")

	hostdata := entity.Sshhostlist{}

	if !utils.CheckNetAddr(host) {
		for {
			host = utils.InputString("请输入正确的主机名(域名或者IP)")
			if utils.CheckNetAddr(host) {
				break
			}
		}
	}
	hostdata.Host = host
	if len(username) <= 0 {
		username = "root"
	}
	hostdata.Username = username
	if hport < 1 || hport > 65535 {
		for {
			hport = utils.InputPort("请输入正确的端口[默认:22]")
			if hport >= 1 && hport <= 65535 {
				break
			}
		}
	}
	hostdata.Port = strconv.Itoa(hport)
	if len(hostdesc) <= 0 {
		for {
			hostdesc = utils.InputString("请输入主机描述")
			if len(hostdesc) > 0 {
				break
			}
		}
	}

	hostdata.Hostdesc = hostdesc

	if iskeyok == "Y" || iskeyok == "yes" || iskeyok == "y" {
		RenderKeyList()
		kid := utils.InputInt("请输入ssh密钥对ID")
		// keyObj, keyObjErr := QueryKeyOneById(kid)
		keyObj, keyObjErr := keylist.QueryOne(kid)
		if keyObjErr != nil {
			RenderKeyList()
			kid = utils.InputInt("请输入正确的ssh密钥对ID")
			keyObj, keyObjErr = keylist.QueryOne(kid)
			if keyObjErr != nil {
				fmt.Println("输入的ssh密钥对ID不正确，请重新操作")
				os.Exit(1)
			}
		}
		hostdata.Keypath = strconv.Itoa(keyObj.Id)
	}
	if initdb.CheckDBIsWritableBool() {
		rs := hostlist.Insert(&hostdata)
		return &hostdata, rs
	} else {
		// 当数据库不可写时候，采用网络方式
		rs := utils.RemoteRequest("/api/hosts/insert", hostdata, "POST")
		return &hostdata, rs
	}
}

func UpdateHost(id int) (*entity.Sshhostlist, bool) {
	data, ObjErr := hostlist.QueryOne(id)
	if ObjErr != nil {
		fmt.Println("输入的主机ID不正确，请重新操作")
		os.Exit(1)
	}
	host := utils.InputString("请输入新主机名(域名或者IP)[" + data.Host + "]")
	if utils.CheckNetAddr(host) {
		data.Host = host
	}

	username := utils.InputString("请输入用户名[" + data.Username + "]")
	if len(username) > 0 {
		data.Username = username
	}
	hport := utils.InputPort("请输入端口[" + data.Port + "]")
	if hport >= 1 && hport <= 65535 {
		data.Port = strconv.Itoa(hport)
	}
	hostdesc := utils.InputString("请输入主机描述[" + data.Hostdesc + "]")
	if len(hostdesc) > 0 {
		data.Hostdesc = hostdesc
	}
	var isKeyDefault string
	if len(data.Keypath) > 0 {
		isKeyDefault = "Y"
	} else {
		isKeyDefault = "N"
	}
	iskeyok := utils.InputString("是否SSH密钥对登录[" + isKeyDefault + "]?[Y/N]")
	if iskeyok == "Y" || iskeyok == "yes" || iskeyok == "y" {
		RenderKeyList()
		kid := utils.InputInt("请输入ssh密钥对ID")
		// keyObj, keyObjErr := QueryKeyOneById(kid)
		keyObj, keyObjErr := keylist.QueryOne(kid)
		if keyObjErr != nil {
			RenderKeyList()
			kid = utils.InputInt("请输入正确的ssh密钥对ID")
			keyObj, keyObjErr = keylist.QueryOne(kid)
			if keyObjErr != nil {
				fmt.Println("输入的ssh密钥对ID不正确，请重新操作")
				os.Exit(1)
			}
		}
		data.Keypath = strconv.Itoa(keyObj.Id)
	}
	if initdb.CheckDBIsWritableBool() {
		rs := hostlist.Update(data)
		return data, rs
	} else {
		rs := utils.RemoteRequest("/api/hosts/update", data, "POST")
		return data, rs
	}
}

func QueryHostOneById(id int) (*entity.Sshhostlist, error) {
	return hostlist.QueryOne(id)
}
func DeleteHostById(id int) bool {
	return hostlist.Delete(id)
}

func RenderKeyList() {
	list, err := keylist.QueryKeylist()
	errors.CheckError(err)
	t := table.NewWriter()
	header := table.Row{"ID", "ssh identity_file"}
	t.AppendHeader(header)
	var rows []table.Row
	for i := 0; i < len(list); i++ {
		item := list[i]
		rows = append(rows, table.Row{item.Id, item.Keyname})
	}
	t.AppendRows(rows)
	fmt.Println(t.Render())
}

func QueryKeyOneById(id int) (*entity.Sshkeylist, error) {
	return keylist.QueryOne(id)
}
func DeleteKeyById(id int, commandLine bool) (bool, string) {
	if commandLine && !initdb.CheckDBIsWritableBool() {
		params := map[string]int{
			"id": id,
		}
		rs := utils.RemoteRequest("/api/keys/insert", params, "DELETE")
		return rs, ""
	} else {
		return keylist.Delete(id)
	}
}

func AddKeyInfoLocal() bool {
	keyPath := utils.InputString("请输入ssh密钥对私钥路径[如: ~/.ssh/id_rsa]")

	if len(keyPath) > 0 && !utils.FileExists(keyPath) {
		fmt.Printf("私钥文件[%s]不存在或者没有可读权限", keyPath)
		os.Exit(1)
	} else if len(keyPath) <= 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("文件[%s]不存在或者没有可读权限", keyPath)
			os.Exit(1)
		} else {
			keyPath = homeDir + "/.ssh/id_rsa"
		}
	}
	pubKeyPath := keyPath + ".pub"
	if !utils.FileExists(pubKeyPath) {
		fmt.Printf("公钥文件[%s]不存在或者没有可读权限", keyPath)
		os.Exit(1)
	}
	data := entity.Sshkeylist{}

	privateContent, err := os.ReadFile(keyPath)
	errors.CheckError(err)
	data.Privatekey = string(privateContent)

	publicContent, err := os.ReadFile(pubKeyPath)
	errors.CheckError(err)
	data.Publickey = string(publicContent)

	fileName := path.Base(keyPath)
	data.Keyname = fileName
	if initdb.CheckDBIsWritableBool() {
		return keylist.Insert(&data)
	} else {
		rs := utils.RemoteRequest("/api/keys/insert", data, "POST")
		return rs
	}
}
func AddKeyInfo() bool {
	keyName := utils.InputString("请输入ssh密钥名称或描述[如: id_rsa]:")
	if len(keyName) <= 0 {
		keyName = "id_rsa"
	}
	keyPasswd := utils.InputString("请输入ssh密钥密码[可以为空]:")
	if len(keyPasswd) <= 0 {
		keyPasswd = ""
	}
	privateKey, publicKey, err := utils.Sshkeygen(keyPasswd)

	if err == nil {
		insertData := &entity.Sshkeylist{}
		insertData.Keyname = keyName
		insertData.Privatekey = privateKey
		insertData.Publickey = publicKey
		if initdb.CheckDBIsWritableBool() {
			return keylist.Insert(insertData)
		} else {
			return utils.RemoteRequest("/api/keys/insert", insertData, "POST")
		}
	} else {
		fmt.Println("生成密钥失败：", err.Error())
		return false
	}
}

func QueryKeysList(page int, limit int) datavo.KeysResp {
	list, total, err := keylist.QueryKeylistPage(page, limit)
	errors.CheckError(err)
	res := datavo.KeysResp{
		Code:  200,
		Msg:   "获取成功",
		Data:  list,
		Count: total,
	}
	if err != nil {
		res.Code = 500
		res.Msg = "获取失败"
	}
	return res
}
func SaveKeys(hostdata *entity.Sshkeylist, izInsert bool) datavo.HostResp {
	var rs bool
	res := datavo.HostResp{
		Code:  200,
		Msg:   "保存成功",
		Count: 1,
	}
	if izInsert {
		rs = keylist.Insert(hostdata)
	} else {
		if hostdata.Id == 0 {
			res.Code = 500
			res.Msg = "保存失败，参数异常，ID不能为空"
			res.Count = 0
			return res
		}
		rs = keylist.Update(hostdata)
	}

	if !rs {
		res.Code = 500
		res.Msg = "保存失败"
		res.Count = 0
	}
	return res
}
func CreateKeys(data *datavo.CreateSshKeyParams) datavo.HostResp {
	var rs bool
	res := datavo.HostResp{
		Code:  200,
		Msg:   "保存成功",
		Count: 1,
	}
	fmt.Println("ck:", data.Passwd)
	privateKey, publicKey, err := utils.Sshkeygen(data.Passwd)

	if err == nil {
		insertData := &entity.Sshkeylist{}
		insertData.Keyname = data.Keyname
		insertData.Privatekey = privateKey
		insertData.Publickey = publicKey
		rs = keylist.Insert(insertData)
	} else {
		res.Code = 500
		res.Msg = err.Error()
		res.Count = 0
		return res
	}

	if !rs {
		res.Code = 500
		res.Msg = "保存失败"
		res.Count = 0
	}
	return res
}

func HomeCount() any {
	hostQty := hostlist.CountTotal()
	keysQty := keylist.CountTotal()
	heQty := hostextent.CountTotal()
	var homeCountResp = datavo.HomeCountResp{
		TotalCount: hostQty,
		KeysCount:  keysQty,
		CloudCount: heQty,
		LocalCount: hostQty - heQty,
	}
	res := datavo.UserResp[datavo.HomeCountResp]{
		Code: 200,
		Msg:  "获取成功",
		Data: homeCountResp,
	}
	return res
}

func QueryHostExtentList(page int, limit int) datavo.PubResp[entity.HostExtent] {
	list, total, err := hostextent.Querylist(page, limit, -1)
	errors.CheckError(err)
	res := datavo.PubResp[entity.HostExtent]{
		Code:  200,
		Msg:   "获取成功",
		Data:  list,
		Count: total,
	}
	if err != nil {
		res.Code = 500
		res.Msg = "获取失败"
	}
	return res
}
func SaveHostExtent(hostdata *entity.HostExtent, izInsert bool) datavo.HostResp {
	var rs bool
	res := datavo.HostResp{
		Code:  200,
		Msg:   "保存成功",
		Count: 1,
	}
	hostobj, herr := hostlist.QueryOne(hostdata.Id)
	if herr != nil {
		res.Code = 500
		res.Msg = "获取主机异常，请重试"
		res.Count = 0
		return res
	}
	hostdata.Host = hostobj.Host + hostobj.Port
	if izInsert {
		rs = hostextent.Insert(hostdata)
	} else {
		if hostdata.Id == 0 {
			res.Code = 500
			res.Msg = "保存失败，参数异常，ID不能为空"
			res.Count = 0
			return res
		}
		rs = hostextent.Update(hostdata)
	}

	if !rs {
		res.Code = 500
		res.Msg = "保存失败"
		res.Count = 0
	}
	return res
}

func DeleteHostExtentById(id int, commandLine bool) (bool, string) {
	if commandLine && !initdb.CheckDBIsWritableBool() {
		params := map[string]int{
			"id": id,
		}
		rs := utils.RemoteRequest("/api/hostExtent/insert", params, "DELETE")
		return rs, ""
	} else {
		return hostextent.Delete(id), ""
	}
}

func SaveUserInfo(hostdata *entity.UserInfo, izInsert bool) datavo.HostResp {
	var rs bool
	res := datavo.HostResp{
		Code:  200,
		Msg:   "保存成功",
		Count: 1,
	}
	if izInsert {
		rs = userinfolist.Insert(hostdata)
	} else {
		if hostdata.Id == 0 {
			res.Code = 500
			res.Msg = "保存失败，参数异常，ID不能为空"
			res.Count = 0
			return res
		}
		rs = userinfolist.Update(hostdata)
	}

	if !rs {
		res.Code = 500
		res.Msg = "保存失败"
		res.Count = 0
	}
	return res
}
func UserLoginCheck(loginParams datavo.LoginParams) (int, string, entity.UserInfo) {
	var rs bool
	userInfo, err := userinfolist.FetchByAccount(loginParams.UserName)
	if err != nil {
		return 500, "用户名或密码错误", entity.UserInfo{}
	}
	rs = sm3utils.VerifySaltedHash(loginParams.Passwd, userInfo.Passwd)
	if !rs {
		return 500, "登录失败，用户名或密码错误", entity.UserInfo{}
	}
	return 200, "登录失败，用户名或密码错误", *userInfo
}

func QueryUserInfoOneById(id int) (*entity.UserInfo, error) {
	return userinfolist.QueryOne(id)
}
