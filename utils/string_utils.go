package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	CHAR_EMPTY          = ""
	CHAR_WHITESPACE     = " "
	CHAR_FORWARD_SLASH  = "/"
	CHAR_BACKWARD_SLASH = "\\"
	CHAR_HYPHEN         = "-"
	CHAR_COMMA          = ","

	BYTE_HYPHEN = '-'
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

}

func ToTitle(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func ToUppper(str string) string {
	return strings.ToUpper(str)
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func RemoveWhiteSpaces(str string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(str, " ")
}

func TrimAll(str string) string {
	return Trim(RemoveWhiteSpaces(str))
}

func RemoveSpecialChars(str string) string {
	str = ReplaceString(str, CHAR_COMMA, CHAR_WHITESPACE)
	str = ReplaceString(str, CHAR_HYPHEN, CHAR_WHITESPACE)
	str = ReplaceString(str, CHAR_FORWARD_SLASH, CHAR_WHITESPACE)
	str = ReplaceString(str, CHAR_BACKWARD_SLASH, CHAR_WHITESPACE)
	return str
}

func ReplaceString(str string, old string, new string) string {
	return strings.ReplaceAll(str, old, new)
}

func IsEmpty(str string) bool {
	return len(Trim(str)) == 0
}

func GenerateKeyIdFromName(name string) string {
	var seoName string = ""

	seoName = ToLower(name) // Convert LowerCase
	seoName = Trim(seoName) // Remove leading and triling spaces
	seoName = RemoveSpecialChars(seoName)
	seoName = RemoveWhiteSpaces(seoName) // Remove multiple spaces in-between the string
	// Finally replace "-" with " "
	seoName = ReplaceString(seoName, CHAR_WHITESPACE, CHAR_HYPHEN)

	return seoName
}

func GenerateNextKeyId(name string) string {
	// Find the LastIndex of '-' character
	i := strings.LastIndexByte(name, BYTE_HYPHEN)
	if i < 0 {
		return name + "-1"
	}

	// Convert it to int
	d, err := strconv.Atoi(name[i+1:])
	if err != nil {
		return name + "-1"
	}

	// Increment the value
	d++
	// Append it to String
	name = name[:i+1] + strconv.Itoa(d)

	// return new string
	return name
}
