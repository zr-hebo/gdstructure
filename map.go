package gdstructure

import "errors"

// GetMapStringKeys 获取golang中string为key的简单map的所有key
func GetMapStringKeys(rawMap interface{}) (keys []string, err error) {
	switch realMap := rawMap.(type) {
	case map[string]bool:
		keys = getBoolMapKeys(realMap)
	case map[string]int:
		keys = getIntMapKeys(realMap)
	case map[string]int8:
		keys = getInt8MapKeys(realMap)
	case map[string]int16:
		keys = getInt16MapKeys(realMap)
	case map[string]int32:
		keys = getInt32MapKeys(realMap)
	case map[string]int64:
		keys = getInt64MapKeys(realMap)
	case map[string]uint:
		keys = getUintMapKeys(realMap)
	case map[string]uint8:
		keys = getUint8MapKeys(realMap)
	case map[string]uint16:
		keys = getUint16MapKeys(realMap)
	case map[string]uint32:
		keys = getUint32MapKeys(realMap)
	case map[string]uint64:
		keys = getUint64MapKeys(realMap)
	case map[string]string:
		keys = getStringMapKeys(realMap)
	case map[string]float32:
		keys = getFloat32MapKeys(realMap)
	case map[string]float64:
		keys = getFloat64MapKeys(realMap)
	default:
		err = errors.New("cannot resolve param to simple string key map")
	}

	return
}

func getUint8MapKeys(rawMap map[string]uint8) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getUint16MapKeys(rawMap map[string]uint16) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getUint32MapKeys(rawMap map[string]uint32) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getUint64MapKeys(rawMap map[string]uint64) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getUintMapKeys(rawMap map[string]uint) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getInt8MapKeys(rawMap map[string]int8) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getInt16MapKeys(rawMap map[string]int16) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getInt32MapKeys(rawMap map[string]int32) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getInt64MapKeys(rawMap map[string]int64) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getIntMapKeys(rawMap map[string]int) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getBoolMapKeys(rawMap map[string]bool) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getStringMapKeys(rawMap map[string]string) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getFloat32MapKeys(rawMap map[string]float32) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}

func getFloat64MapKeys(rawMap map[string]float64) (keys []string) {
	keys = make([]string, 0, len(rawMap))
	for key := range rawMap {
		keys = append(keys, key)
	}

	return
}
