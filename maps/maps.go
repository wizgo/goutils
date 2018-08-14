package maps

type Map interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat(key string) float64
	GetString(key string) string
}

func NewMap() Map {
	mp := StringMap(make(map[string]interface{}))
	return &mp
}

func NewStringMap() Map {
	mp := StringMap(make(map[string]interface{}))
	return &mp
}

func NewGenericMap() Map {
	mp := GenericMap(make(map[interface{}]interface{}))
	return &mp
}

func FromStringMap(m map[string]interface{}) Map {
	mp := StringMap(m)
	return &mp
}

func FromGenericMap(m map[interface{}]interface{}) Map {
	mp := GenericMap(m)
	return &mp
}

func FromMap(v interface{}) Map {
	if v != nil {
		if mv, ok := v.(map[string]interface{}); ok {
			m := StringMap(mv)
			return &m
		}
		if mv, ok := v.(map[interface{}]interface{}); ok {
			m := GenericMap(mv)
			return &m
		}
	}
	return nil
}
