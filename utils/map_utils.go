package utils

type Map map[string]interface{}

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
