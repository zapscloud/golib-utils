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
