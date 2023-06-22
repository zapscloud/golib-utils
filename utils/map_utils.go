package utils

type Map map[string]interface{}

func GetMemberDataStr(data Map, memberName string) (string, error) {

	dataVal, dataOk := data[memberName]

	if !dataOk || IsEmpty(dataVal.(string)) {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data/Empty", ErrorDetail: memberName + " data must be sent or should not be empty"}
		return "", err
	}

	return dataVal.(string), nil
}

func GetMemberDataInt(data Map, memberName string) (int, error) {

	dataVal, dataOk := data[memberName]

	if !dataOk {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " data must be sent"}
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
