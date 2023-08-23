package utils

type Map map[string]interface{}

func GetMemberData(data Map, memberName string) (any, error) {

	var err error = nil

	// Check datamember
	dataVal, dataOk := data[memberName]

	if !dataOk {
		err = &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
	}

	return dataVal, err
}

func GetMemberDataStr(data Map, memberName string) (string, error) {

	var strReturn string = ""
	var err error = nil
	var isString bool

	// Check datamember
	dataVal, dataOk := data[memberName]

	if dataOk {
		strReturn, isString = dataVal.(string)
		if !isString {
			err = &AppError{ErrorStatus: 400, ErrorMsg: "Invalid Datatype", ErrorDetail: memberName + " value should be a string"}
		} else if IsEmpty(strReturn) {
			err = &AppError{ErrorStatus: 400, ErrorMsg: "Data Empty", ErrorDetail: memberName + " value should not be empty"}
		}
	} else {
		err = &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
	}

	return strReturn, err
}

func GetMemberDataInt(data Map, memberName string) (int, error) {

	var intReturn int = 0
	var err error = nil
	var isInt bool

	// Check datamember
	dataVal, dataOk := data[memberName]

	if dataOk {
		intReturn, isInt = dataVal.(int)
		if !isInt {
			err = &AppError{ErrorStatus: 400, ErrorMsg: "Invalid Datatype", ErrorDetail: memberName + " value should be a integer"}
		}
	} else {
		err = &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
	}

	return intReturn, err
}

func CopyMap(src Map) Map {

	dst := Map{}
	for key, value := range src {
		dst[key] = value
	}
	return dst

}

func MergeMap(dstMap Map, srcMap Map, createNew bool) Map {

	var newMap Map

	if createNew {
		newMap = CopyMap(dstMap)
	} else {
		newMap = dstMap
	}

	// Enumerate srcMap and append it to dstMap
	for key, value := range srcMap {
		newMap[key] = value
	}

	return newMap
}
