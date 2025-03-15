package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
)

func EncryptPassword(data []byte) (res string) {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 密码的加盐存储和匹配

const SaltSize = 16

func RandomSalt() []byte {
	return generateRandomSalt(SaltSize)
}

func generateRandomSalt(saltSize int) []byte {
	r := rand.New(rand.NewSource(NowMilliSecond()))
	var salt = make([]byte, saltSize)
	_, err := r.Read(salt)
	// _, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return salt
}

func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}

func DoPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}
