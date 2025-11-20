package password

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GenerateSalt 生成随机盐值
// GenerateSalt generates a random salt
func GenerateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// HashPasswordWithSalt 使用盐值对密码进行哈希
// HashPasswordWithSalt hashes password with salt
func HashPasswordWithSalt(password, salt string) string {
	// 将密码和盐值组合后进行哈希
	// Combine password and salt then hash
	combined := password + salt
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

// VerifyPassword 验证密码是否正确
// VerifyPassword verifies if the password is correct
func VerifyPassword(password, salt, storedHash string) bool {
	computedHash := HashPasswordWithSalt(password, salt)
	return computedHash == storedHash
}

// GeneratePasswordHash 生成密码哈希和盐值
// GeneratePasswordHash generates password hash and salt
func GeneratePasswordHash(password string) (string, string, error) {
	salt, err := GenerateSalt(16) // 16字节盐值 / 16 bytes salt
	if err != nil {
		return "", "", fmt.Errorf("生成盐值失败 / failed to generate salt: %v", err)
	}
	hash := HashPasswordWithSalt(password, salt)
	return hash, salt, nil
}