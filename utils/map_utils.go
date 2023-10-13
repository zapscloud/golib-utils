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

func GetMemberDataInt(data Map, memberName string, ignoreFloat bool) (int, error) {

	var retVal int = 0
	var err error = nil

	// Check datamember
	dataVal, dataOk := data[memberName]

	if dataOk {
		if IsTypeInt(dataVal) {
			retVal = dataVal.(int)
		} else if IsTypeInt32(dataVal) {
			retVal = int(dataVal.(int32))
		} else if IsTypeInt64(dataVal) {
			retVal = int(dataVal.(int64))
		} else if ignoreFloat {
			if IsTypeFloat32(dataVal) {
				retVal = int(dataVal.(float32))
			} else if IsTypeFloat64(dataVal) {
				retVal = int(dataVal.(float64))
			}
		} else {
			err = &AppError{ErrorStatus: 400, ErrorMsg: "Invalid Datatype", ErrorDetail: memberName + " value should be a integer"}
		}
	} else {
		err = &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
	}

	return retVal, err
}

func GetMemberDataBool(data Map, memberName string) (bool, error) {

	// Check datamember
	dataVal, dataOk := data[memberName]

	if dataOk {
		if !IsTypeBool(dataVal) {
			err := &AppError{ErrorStatus: 400, ErrorMsg: "Invalid Datatype", ErrorDetail: memberName + " value should be a Boolean"}
			return false, err
		}
	} else {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: memberName + " value should be sent"}
		return false, err
	}

	return dataVal.(bool), nil
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
