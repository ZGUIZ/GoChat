package utils

import (
	"reflect"
	"strconv"
)

func interfaceToMapByReflection(v interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(v)

	keys := val.MapKeys()
	for i := 0; i < len(keys); i++ {
		var fieldName = keys[i].String()
		fieldValue := val.MapIndex(keys[i]).Interface()
		result[fieldName] = fieldValue
	}
	return result
}

func interfaceToString(v interface{}) string {
	//int转string
	i, err := v.(int)
	if err {
		return strconv.Itoa(i)
	}
	s, err := v.(string)
	if !err {
		return ""
	}
	return s
}
