package jwtutils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hlinfocc/cySSHClient2/pkg/datavo"
)

// 密钥（实际应用中应从安全配置获取）
var jwtSecret = []byte("your_secret_key")

func GenJwtToken(id int, realName string, account string, status int, userType int, role string) (string, error) {
	// 生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datavo.CustomClaims{
		UserID:   id, // 用户ID
		RealName: realName,
		Account:  account,
		Status:   status,
		UserType: userType,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			Issuer:    "ylcxy.cn",                                         // 签发者
		},
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JWTParse(tokenString string) (datavo.CustomClaims, error) {
	// 验证并解析token
	token, err := jwt.ParseWithClaims(tokenString, &datavo.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return datavo.CustomClaims{}, fmt.Errorf("无效的令牌")
	}
	claims, ok := token.Claims.(*datavo.CustomClaims)
	if ok {
		return *claims, nil
	}
	return datavo.CustomClaims{}, fmt.Errorf("令牌解析失败")
}
