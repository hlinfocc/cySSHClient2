package datavo

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
)

type SimpResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type HostResp struct {
	Code  int                   `json:"code"`
	Msg   string                `json:"msg"`
	Data  []*entity.HostlistAll `json:"data"`
	Count int64                 `json:"count"`
}

type KeysResp struct {
	Code  int                  `json:"code"`
	Msg   string               `json:"msg"`
	Data  []*entity.Sshkeylist `json:"data"`
	Count int64                `json:"count"`
}

type UserResp[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
	Count int64  `json:"count"`
}

type HostParams struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	Description string `json:"description"`
	HostIp      string `json:"hostip"`
	HostExtent  int    `json:hostExtent`
}

type TokenInfo struct {
	Token    string          `json:"token"`
	UserInfo entity.UserInfo `json:"userInfo"`
}

type CreateSshKeyParams struct {
	Keyname string `json:"keyname"`
	Passwd  string `json:"passwd"`
}

type HomeCountResp struct {
	TotalCount int64 `json:"totalCount"`
	CloudCount int64 `json:"cloudCount"`
	LocalCount int64 `json:"localCount"`
	KeysCount  int64 `json:"keysCount"`
}

type PubResp[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  []*T   `json:"data"`
	Count int64  `json:"count"`
}

type LoginParams struct {
	UserName string `json:"account"`
	Passwd   string `json:"password"`
}

// 定义Claims结构体
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	RealName string `json:"realName"`
	Account  string `json:"account"`
	Status   int    `json:"status"`
	UserType int    `json:"userType"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
