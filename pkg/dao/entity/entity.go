package entity

type Sshhostlist struct {
	Id       int    `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null" json:"id"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Port     string `json:"port"`
	Iskey    int    `json:"iskey"`
	Keypath  string `json:"keypath"`
	Hostdesc string `json:"hostdesc"`
}

type Sshkeylist struct {
	Id         int    `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null" json:"id"`
	Keyname    string `json:"keyname"`
	Privatekey string `json:"privatekey"`
	Publickey  string `json:"publickey"`
}

type HostlistAll struct {
	Id       int    `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null" json:"id"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Port     string `json:"port"`
	Iskey    int    `json:"iskey"`
	Keypath  string `json:"keypath"`
	Hostdesc string `json:"hostdesc"`
	Keyname  string `json:"keyname"`
}

type UserInfo struct {
	Id            int    `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null" json:"id"`
	RealName      string `json:"realName"`
	Account       string `json:"account"`
	Passwd        string `gorm:"column:passwd" json:"-"` // 序列化忽略但允许反序列化
	Status        int    `json:"status"`
	UserType      int    `json:"userType"`
	Role          string `json:"role"`
	LastLoginTime string `json:"lastLoginTime"`
	LastLoginIp   string `json:"lastLoginIp"`
	ThatLoginTime string `json:"thatLoginTime"`
	ThatLoginIp   string `json:"thatLoginIp"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
	Isdelete      int    `json:"isdelete"`
}

type HostExtent struct {
	Id        int    `gorm:"column:id;PRIMARY_KEY;not null" json:"id"`
	Host      string `json:"host"`
	CloudType string `json:"cloudType"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	IzCrond   int    `json:"izCrond"` //是否开启定时任务计划，0 不开启，1开启
	Remarks   string `json:"remarks"`
}
