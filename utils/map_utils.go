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

	// Check datamember
	dataVal, dataOk := data[memberName]

	if dataOk {
		if !IsTypeString(dataVal) {
			err := &AppError{ErrorStatus: 400, ErrorMsg: "Invalid Datatype", ErrorDetail: memberName + " value should be a string"}
			return "", err
		} else if IsEmpty(dataVal.(string)) {
			err := &AppError{ErrorStatus: 400, ErrorMsg: "Data Empty", ErrorDetail: memberName + " value should not be empty"}
			return "", err
		}
	} else {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
		return "", err
	}

	return dataVal.(string), nil
}

func GetMemberDataInt(data Map, memberName string) (int, error) {

	// Check datamember
	dataVal, dataOk := data[memberName]

	if dataOk {
		if !IsTypeInt(dataVal) {
			err := &AppError{ErrorStatus: 400, ErrorMsg: "Invalid Datatype", ErrorDetail: memberName + " value should be a integer"}
			return 0, err
		}
	} else {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
		return 0, err
	}

	return dataVal.(int), nil
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
