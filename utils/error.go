package utils

import (
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
	//return fmt.Sprintf("%s - %s", e.ErrorCode, e.ErrorMsg)
	var err string = "Error"

	if len(e.ErrorDetail) > 0 {
		err = e.ErrorDetail
	} else {
		if len(e.ErrorCode) > 0 {
			err = e.ErrorCode + " - " + e.ErrorMsg
		} else if len(e.ErrorMsg) > 0 {
			err = e.ErrorMsg
		}
	}

	return err
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}
