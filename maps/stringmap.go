package maps

import (
	"github.com/wizgo/goutils/cast"
)

type StringMap map[string]interface{}

// public interfaces
func (m *StringMap) Set(key string, value interface{}) {
	_Set(m, key, value)
}

func (m *StringMap) Get(key string) interface{} {
	return _Get(m, key)
}

func (m *StringMap) GetBool(key string) (b bool) {
	if v := m.Get(key); v != nil {
		b = cast.CastBool(v)
	}
	return
}

func (m *StringMap) GetInt(key string) (n int) {
	if v := m.Get(key); v != nil {
		n = cast.CastInt(v)
	}
	return
}

func (m *StringMap) GetFloat(key string) (f float64) {
	if v := m.Get(key); v != nil {
		f = cast.CastFloat(v)
	}
	return
}

func (m *StringMap) GetString(key string) (s string) {
	if v := m.Get(key); v != nil {
		s = cast.CastString(v)
	}
	return
}

// private interfaces
func (m *StringMap) _Set(key string, value interface{}) {
	(*m)[key] = value
}

func (m *StringMap) _Get(key string) interface{} {
	return (*m)[key]
}

func (m *StringMap) _InitSubMap(key string) {
	m._Set(key, make(map[string]interface{}))
}
