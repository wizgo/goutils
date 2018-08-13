package cast

import (
	"encoding/json"
	"strconv"
)

func CastBool(v interface{}) (b bool) {
	switch s := v.(type) {
	case bool:
		b = s
	case string:
		b, _ = strconv.ParseBool(s)
	case []byte:
		b, _ = strconv.ParseBool(string(s))
	case int:
		b = s != 0
	case int8:
		b = s != 0
	case int16:
		b = s != 0
	case int32:
		b = s != 0
	case int64:
		b = s != 0
	case uint:
		b = s != 0
	case uint8:
		b = s != 0
	case uint16:
		b = s != 0
	case uint32:
		b = s != 0
	case uint64:
		b = s != 0
	case float32:
		b = s != float32(0)
	case float64:
		b = s != float64(0)
	}
	return
}

func CastInt(v interface{}) (n int) {
	switch s := v.(type) {
	case bool:
		if s {
			n = 1
		}
	case string:
		n, _ = strconv.Atoi(s)
	case []byte:
		n, _ = strconv.Atoi(string(s))
	case int:
		n = s
	case int8:
		n = int(s)
	case int16:
		n = int(s)
	case int32:
		n = int(s)
	case int64:
		n = int(s)
	case uint:
		n = int(s)
	case uint8:
		n = int(s)
	case uint16:
		n = int(s)
	case uint32:
		n = int(s)
	case uint64:
		n = int(s)
	case float32:
		n = int(s)
	case float64:
		n = int(s)
	}
	return
}

func CastFloat(v interface{}) (f float64) {
	switch s := v.(type) {
	case bool:
		if s {
			f = 1.0
		}
	case string:
		f, _ = strconv.ParseFloat(s, 64)
	case []byte:
		f, _ = strconv.ParseFloat(string(s), 64)
	case int:
		f = float64(s)
	case int8:
		f = float64(s)
	case int16:
		f = float64(s)
	case int32:
		f = float64(s)
	case int64:
		f = float64(s)
	case uint:
		f = float64(s)
	case uint8:
		f = float64(s)
	case uint16:
		f = float64(s)
	case uint32:
		f = float64(s)
	case uint64:
		f = float64(s)
	case float32:
		f = float64(s)
	case float64:
		f = s
	}
	return
}

func CastString(v interface{}) (str string) {
	switch s := v.(type) {
	case string:
		str = s
	case bool:
		str = strconv.FormatBool(s)
	case []byte:
		str = string(s)
	case int:
		str = strconv.FormatInt(int64(s), 10)
	case int8:
		str = strconv.FormatInt(int64(s), 10)
	case int16:
		str = strconv.FormatInt(int64(s), 10)
	case int32:
		str = strconv.FormatInt(int64(s), 10)
	case int64:
		str = strconv.FormatInt(s, 10)
	case uint:
		str = strconv.FormatUint(uint64(s), 10)
	case uint8:
		str = strconv.FormatUint(uint64(s), 10)
	case uint16:
		str = strconv.FormatUint(uint64(s), 10)
	case uint32:
		str = strconv.FormatUint(uint64(s), 10)
	case uint64:
		str = strconv.FormatUint(s, 10)
	case float32:
		str = strconv.FormatFloat(float64(s), 'f', -1, 32)
	case float64:
		str = strconv.FormatFloat(s, 'f', -1, 64)
	default:
		b, _ := json.Marshal(v)
		str = string(b)
	}
	return
}
