package maps

import (
	"reflect"
	"strconv"
	"strings"
)

type _map_ interface {
	_Get(key string) interface{}
	_Set(key string, value interface{})
	_InitSubMap(key string)
}

func _Set(m _map_, key string, value interface{}) {
	if strings.Contains(key, ".") {
		k1, k2 := _SplitKey(key)
		if sub := FromMap(m._Get(k1)); sub != nil {
			sub.Set(k2, value)
		} else {
			m._InitSubMap(k1)
			FromMap(m._Get(k1)).Set(k2, value)
		}
	} else {
		m._Set(key, value)
	}
}

func _Get(m _map_, key string) interface{} {
	if strings.Contains(key, ".") {
		k1, k2 := _SplitKey(key)
		if sub := FromMap(_Get(m, k1)); sub != nil {
			return sub.Get(k2)
		}
	} else if k, i := _SplitKeyIndex(key); k != "" {
		if v := m._Get(k); v != nil {
			if av := reflect.ValueOf(v); (av.Kind() == reflect.Slice || av.Kind() == reflect.Array) && av.Len() > i {
				return av.Index(i).Interface()
			}
		}
	} else {
		return m._Get(key)
	}
	return nil
}

func _SplitKey(key string) (string, string) {
	k := strings.SplitN(key, ".", 2)
	return k[0], k[1]
}

func _SplitKeyIndex(key string) (k string, i int) {
	if ln, pos := len(key), strings.Index(key, "["); pos > 0 && pos < ln-1 && key[ln-1:] == "]" {
		if index, err := strconv.Atoi(key[pos+1 : ln-1]); err == nil {
			k, i = key[:pos], index
		}
	}
	return
}
