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
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateChecksumId(prefix string, key string) string {
	hash := md5.Sum([]byte(key))
	new_id := fmt.Sprintf("%s_%x", prefix, hash)
	return new_id
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

func IsTypeInt32(data any) bool {
	_, isInt := data.(int32)
	return isInt
}

func IsTypeInt64(data any) bool {
	_, isInt := data.(int64)
	return isInt
}

func IsTypeFloat32(data any) bool {
	_, isFloat := data.(float32)
	return isFloat
}

func IsTypeFloat64(data any) bool {
	_, isFloat := data.(float64)
	return isFloat
}

func IsTypeBool(data any) bool {
	_, isBool := data.(bool)
	return isBool
}

func IsTypeMap(data any) bool {
	_, isMap := data.(Map)
	return isMap
}

func IsTypeMapStrInterface(data any) bool {
	_, isMap := data.(map[string]interface{})
	return isMap
}
