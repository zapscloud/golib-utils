package utils

import (
	"log"
)

// AppFunction -- Application Function Structure
type AppFunction struct {
	FunctionName string `json:"function_name"`
	Permission   string `json:"permission"`
}

type AppSubModule struct {
	SubModuleName string        `json:"module_name"`
	Permission    string        `json:"permission"`
	Functions     []AppFunction `json:"functions"`
}

type AppModule struct {
	ModuleName string         `json:"module_name"`
	Permission string         `json:"permission"`
	SubModules []AppSubModule `json:"submodules"`
}

type AppActions struct {
	ApplicationName string      `json:"app_name"`
	Permission      string      `json:"permission"`
	Modules         []AppModule `json:"modules"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func IsMemberExist(data Map, member string) (string, error) {

	memberId, dataok := data[member]

	if !dataok {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: member + " data must be sent"}
		return "", err
	}

	return memberId.(string), nil
}
