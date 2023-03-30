package utils

import (
	"fmt"
	"log"
)

// Error -- Application Error Structure
type AppError struct {
	ErrorStatus int
	ErrorCode   string /* [A..action/S.. Service/D.. Dao /X.. External API][X AABBCCEE]*/
	ErrorMsg    string
	ErrorDetail string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrorCode, e.ErrorMsg)
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}
