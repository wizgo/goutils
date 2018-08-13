package maps

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"dzh.com/cast"
	"gopkg.in/yaml.v2"
)

type Map struct {
	M *map[string]interface{}
}

func NewMap() *Map {
	m := make(map[string]interface{})
	return &Map{M: &m}
}

func (m *Map) LoadJson(data []byte) error {
	return json.Unmarshal(data, m.M)
}

func (m *Map) ToJson() (data []byte) {
	data, _ = json.Marshal(m.M)
	return
}

func (m *Map) LoadYaml(data []byte) error {
	return yaml.Unmarshal(data, m.M)
}

func (m *Map) ToYaml() (data []byte) {
	data, _ = yaml.Marshal(m.M)
	return
}

func (m *Map) Set(key string, value interface{}) {
	Set(m.M, key, value)
}

func (m *Map) Get(key string) interface{} {
	return Get(m.M, key)
}

func (m *Map) GetBool(key string) bool {
	return GetBool(m.M, key)
}

func (m *Map) GetInt(key string) int {
	return GetInt(m.M, key)
}

func (m *Map) GetFloat(key string) float64 {
	return GetFloat(m.M, key)
}

func (m *Map) GetString(key string) string {
	return GetString(m.M, key)
}

func Set(mp *map[string]interface{}, key string, value interface{}) {
	if strings.Contains(key, ".") {
		k := strings.SplitN(key, ".", 2)
		if v := Get(mp, k[0]); v != nil {
			if vv := reflect.ValueOf(v); vv.Kind() == reflect.Map {
				if mv, ok := vv.Interface().(map[string]interface{}); ok {
					Set(&mv, k[1], value)
				}
			}
		} else {
			newv := make(map[string]interface{})
			(*mp)[k[0]] = newv
			Set(&newv, k[1], value)
		}
	} else {
		(*mp)[key] = value
	}
}

func Get(mp *map[string]interface{}, key string) interface{} {
	if strings.Contains(key, ".") {
		k := strings.SplitN(key, ".", 2)
		if v := Get(mp, k[0]); v != nil {
			if vv := reflect.ValueOf(v); vv.Kind() == reflect.Map {
				if mv, ok := vv.Interface().(map[string]interface{}); ok {
					return Get(&mv, k[1])
				}
				if mv, ok := vv.Interface().(map[interface{}]interface{}); ok {
					return GenericGet(&mv, k[1])
				}
			}
		}
	} else if k, i := _split_key_index(key); k != "" {
		if v, ok1 := (*mp)[k]; ok1 && v != nil {
			if vv := reflect.ValueOf(v); (vv.Kind() == reflect.Slice || vv.Kind() == reflect.Array) && vv.Len() > i {
				return vv.Index(i).Interface()
			}
		}
	} else {
		return (*mp)[key]
	}
	return nil
}

func GenericGet(mp *map[interface{}]interface{}, key string) interface{} {
	if strings.Contains(key, ".") {
		k := strings.SplitN(key, ".", 2)
		if v := GenericGet(mp, k[0]); v != nil {
			if vv := reflect.ValueOf(v); vv.Kind() == reflect.Map {
				if mv, ok := vv.Interface().(map[interface{}]interface{}); ok {
					return GenericGet(&mv, k[1])
				}
			}
		}
	} else if k, i := _split_key_index(key); k != "" {
		if v, ok1 := (*mp)[k]; ok1 && v != nil {
			if vv := reflect.ValueOf(v); (vv.Kind() == reflect.Slice || vv.Kind() == reflect.Array) && vv.Len() > i {
				return vv.Index(i).Interface()
			}
		}
	} else {
		return (*mp)[key]
	}
	return nil
}

func GetBool(mp *map[string]interface{}, key string) (b bool) {
	if v := Get(mp, key); v != nil {
		b = cast.CastBool(v)
	}
	return
}

func GetInt(mp *map[string]interface{}, key string) (n int) {
	if v := Get(mp, key); v != nil {
		n = cast.CastInt(v)
	}
	return
}

func GetFloat(mp *map[string]interface{}, key string) (f float64) {
	if v := Get(mp, key); v != nil {
		f = cast.CastFloat(v)
	}
	return
}

func GetString(mp *map[string]interface{}, key string) (str string) {
	if v := Get(mp, key); v != nil {
		str = cast.CastString(v)
	}
	return
}

func _split_key_index(key string) (k string, i int) {
	if ln, pos := len(key), strings.Index(key, "["); key[ln-1:] == "]" && pos > 0 && pos < ln-1 {
		if index, err := strconv.Atoi(key[pos+1 : ln-1]); err == nil {
			k, i = key[:pos], index
		}
	}
	return
}
