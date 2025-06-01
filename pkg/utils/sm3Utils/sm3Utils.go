package sm3utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/emmansun/gmsm/sm3"
)

// 生成加盐哈希 (格式: base64(盐) + "|" + base64(哈希值))
func GenerateSaltedHash(data string) string {
	// 1. 生成随机盐值 (16字节)
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "" //, fmt.Errorf("生成盐失败: %v", err)
	}

	// 2. 创建SM3哈希实例
	hash := sm3.New()

	// 3. 先写入盐值再写入数据
	hash.Write(salt)
	hash.Write([]byte(data))

	// 4. 计算哈希值
	hashValue := hash.Sum(nil)

	// 5. 组合盐值和哈希值 (Base64编码)
	encodedSalt := hex.EncodeToString(salt)
	encodedHash := hex.EncodeToString(hashValue)
	// 6. 返回组合字符串 (盐|哈希)
	return fmt.Sprintf("$2y$10$%s.%s", encodedSalt, encodedHash)
}

// 验证加盐哈希
func VerifySaltedHash(data, storedHash string) bool {
	// 1. 分离存储的盐和哈希
	var storedSalt, storedHashValue string
	if !strings.HasPrefix(storedHash, "$2y$10$") {
		return false
	}
	storedHash = strings.ReplaceAll(storedHash, "$2y$10$", "")
	hashArr := strings.Split(storedHash, ".")
	storedSalt = hashArr[0]
	storedHashValue = hashArr[1]
	// 2. 解码十六进制字符串的盐值
	salt, _ := hex.DecodeString(storedSalt)
	// 3. 重新计算哈希
	hash := sm3.New()
	hash.Write(salt)
	hash.Write([]byte(data))
	computedHash := hash.Sum(nil)

	// 4. 解码存储的哈希值并比较
	decodedStoredHash, _ := hex.DecodeString(storedHashValue)

	// 5. 安全比较哈希值
	return secureCompare(decodedStoredHash, computedHash)
}

// 安全比较两个字节切片 (恒定时间比较)
func secureCompare(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

func Test() {
	// 示例用法
	password := "123456"

	// 生成加盐哈希
	saltedHash := GenerateSaltedHash(password)

	fmt.Printf("存储的加盐哈希: %s\n", saltedHash)

	// 验证密码
	valid := VerifySaltedHash(password, saltedHash)
	fmt.Printf("验证结果: %v\n", valid)

	// 错误密码测试
	invalid := VerifySaltedHash("wrongPassword", saltedHash)
	fmt.Printf("错误密码验证: %v\n", invalid)
}
