package entity

type Sshhostlist struct {
	Id       int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Host     string
	Username string
	Port     string
	Iskey    int
	Keypath  string
	Hostdesc string
}

type Sshkeylist struct {
	Id         int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Keyname    string
	Privatekey string
	Publickey  string
}

type HostlistAll struct {
	Id       int `gorm:"column:id;PRIMARY_KEY;autoIncrement;not null"`
	Host     string
	Username string
	Port     string
	Iskey    int
	Keypath  string
	Hostdesc string
	Keyname  string
}
