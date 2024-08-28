package models

import (
	"encoding/json"
	"reflect"
)

func OmitEmpty(raw json.RawMessage) (json.RawMessage, error) {
	var data map[string]any
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	} else {
		StripEmptyEntries(data)
		return json.Marshal(data)
	}
}

func StripEmptyEntries(data map[string]any) {
	for key, value := range data {
		if isEmpty(reflect.ValueOf(value)) {
			delete(data, key)
		} else if nested, ok := value.(map[string]any); ok { // recursively strip nested maps
			StripEmptyEntries(nested)
		} else if slice, ok := value.([]any); ok {
			value = make([]any, len(value.([]any)))
			i := 0
			for _, item := range slice {
				if mapValue, ok := item.(map[string]any); ok {
					StripEmptyEntries(mapValue)
				}
				if !isEmpty(reflect.ValueOf(item)) {
					value.([]any)[i] = item
					i++
				}
			}
			value = value.([]any)[:i]
		}

		// Strip top level if empty post recursive strip
		if _, ok := data[key]; ok && isEmpty(reflect.ValueOf(value)) {
			delete(data, key)
		}
	}
}

func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return value.Bool() == false
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0
	case reflect.Interface, reflect.Pointer:
		return value.IsNil()
	case reflect.Invalid:
		return true
	default:
		return false
	}
}
