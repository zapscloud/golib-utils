package utils

type Map map[string]interface{}

func IsMemberExist(data Map, member string) (string, error) {

	memberId, dataok := data[member]

	if !dataok || IsEmpty(memberId.(string)) {
		err := &AppError{ErrorStatus: 400, ErrorMsg: "Missing Data", ErrorDetail: member + " data must be sent"}
		return "", err
	}

	return memberId.(string), nil
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
