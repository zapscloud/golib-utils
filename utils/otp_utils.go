package utils

import (
	"crypto/rand"
	"io"
)

func GetRandomOTP(maxDigits int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	// Default OTP number
	var strOTP string = "000000"

	b := make([]byte, maxDigits)
	n, _ := io.ReadAtLeast(rand.Reader, b, maxDigits)
	if n == maxDigits {
		for i := 0; i < len(b); i++ {
			b[i] = table[int(b[i])%len(table)]
		}
		// Update with generated OTP number
		strOTP = string(b)
	}
	return strOTP
}
