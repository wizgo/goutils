package maps

import (
	"github.com/wizgo/goutils/cast"
)

type GenericMap map[interface{}]interface{}

// public interfaces
func (m *GenericMap) Set(key string, value interface{}) {
	_Set(m, key, value)
}

func (m *GenericMap) Get(key string) interface{} {
	return _Get(m, key)
}

func (m *GenericMap) GetBool(key string) (b bool) {
	if v := m.Get(key); v != nil {
		b = cast.CastBool(v)
	}
	return
}

func (m *GenericMap) GetInt(key string) (n int) {
	if v := m.Get(key); v != nil {
		n = cast.CastInt(v)
	}
	return
}

func (m *GenericMap) GetFloat(key string) (f float64) {
	if v := m.Get(key); v != nil {
		f = cast.CastFloat(v)
	}
	return
}

func (m *GenericMap) GetString(key string) (s string) {
	if v := m.Get(key); v != nil {
		s = cast.CastString(v)
	}
	return
}

// private interfaces
func (m *GenericMap) _Set(key string, value interface{}) {
	(*m)[key] = value
}

func (m *GenericMap) _Get(key string) interface{} {
	return (*m)[key]
}

func (m *GenericMap) _InitSubMap(key string) {
	m._Set(key, make(map[interface{}]interface{}))
}
