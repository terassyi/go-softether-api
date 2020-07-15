package pkg

import (
	"encoding/json"
	"strings"
	"time"
)

func StructToMap(params Params) (map[string]interface{}, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(data, &dataMap); err != nil {
		return nil, err
	}
	return splitTypes(dataMap), nil
}

// params typ must be empty.
func MapToStruct(data map[string]interface{}, typ Params) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, typ)
}

func splitTypes(params map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	for k, v := range params {
		s := strings.Split(k, "_")
		if len(s) > 2 {
			t := s[0]
			for i := 1; i < len(s)-1; i++ {
				t += "_" + s[i]
			}
			data[t] = v
			continue
		}
		data[s[0]] = v
	}
	return data
}

func ValidateType(typ string, val interface{}) bool {
	switch typ {
	case str, utf, bin, ip:
		switch val.(type) {
		case string:
			return true
		default:
			return false
		}
	case u32, u64:
		switch val.(type) {
		case int:
			return true
		default:
			return false
		}
	case boolean:
		switch val.(type) {
		case bool:
			return true
		default:
			return false
		}
	case date:
		switch val.(type) {
		case time.Time:
			return true
		default:
			return false
		}
	default:
		return false
	}
}
