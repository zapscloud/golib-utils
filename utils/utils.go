package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"regexp"

	"github.com/rs/xid"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

}

// ValidateDocumentKey -
func ValidateDocumentKey(id string) (bool, error) {
	alphacheck := "^[a-zA-Z][A-Za-z0-9_-]+$"
	match, err := regexp.MatchString(alphacheck, id)
	return match, err
}

// func RandomString(length int) string {
// 	rand.Seed(time.Now().UnixNano())
// 	b := make([]byte, length)
// 	rand.Read(b)
// 	return fmt.Sprintf("%x", b)[:length]
// }

func GenerateUniqueId(_prefix string) (prefix string) {
	guid := xid.New()
	prefix = _prefix + "_" + guid.String()
	return
}

func MakeRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func MakeCapitalRandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateChecksumId(prefix string, key string) (new_id string) {
	hash := md5.Sum([]byte(key))

	new_id = fmt.Sprintf("%s_%x", prefix, hash)
	fmt.Println("New Code ", new_id)

	return
}

func SHA(text string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func IsTypeString(data any) bool {
	_, isString := data.(string)
	return isString
}

func IsTypeInt(data any) bool {
	_, isInt := data.(int)
	return isInt
}

func IsTypeFloat64(data any) bool {
	_, isFloat := data.(float64)
	return isFloat
}
