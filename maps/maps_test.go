package maps_test

import (
	"testing"

	"github.com/wizgo/goutils/maps"
	"gopkg.in/yaml.v2"
)

type test_arg struct {
	k string
	v interface{}
}

var (
	ymlstr = `
a:
    b:
        str: a_b_c
        int_array: [1, 2, 3]
        str_array: [s1, s2, s3]
    c:
        -
            name: admin
            level: 1
        -
            name: user
            level: 2
        -
            name: guest
            level: 3
`
	args_set = []test_arg{
		test_arg{"a.b.str", "a_b_c"},
		test_arg{"a.b.int_array", []int{1, 2, 3}},
		test_arg{"a.b.str_array", []string{"s1", "s2", "s3"}},

		test_arg{"a.c", []map[string]interface{}{
			map[string]interface{}{"name": "admin", "level": 1},
			map[string]interface{}{"name": "user", "level": 2},
			map[string]interface{}{"name": "guest", "level": 3},
		}},
	}
	args_expect = []test_arg{
		test_arg{"", nil},
		test_arg{"a.b.str", "a_b_c"},
		test_arg{"a.b.int_array[0]", 1},
		test_arg{"a.b.int_array[1]", 2},
		test_arg{"a.b.int_array[2]", 3},
		test_arg{"a.b.str_array[0]", "s1"},
		test_arg{"a.b.str_array[1]", "s2"},
		test_arg{"a.b.str_array[2]", "s3"},
		test_arg{"a.c[0].name", "admin"},
		test_arg{"a.c[0].level", 1},
		test_arg{"a.c[1].name", "user"},
		test_arg{"a.c[1].level", 2},
		test_arg{"a.c[2].name", "guest"},
		test_arg{"a.c[2].level", 3},
	}
)

func TestYml(t *testing.T) {
	m := maps.NewMap()
	err := yaml.Unmarshal([]byte(ymlstr), m)
	if err != nil {
		t.Fatalf("failed to parse yaml: %s: %v", ymlstr, err)
	}
	for _, arg := range args_expect {
		if v := m.Get(arg.k); v != arg.v {
			t.Fatalf("for key '%s', expect '%v', but was '%v'", arg.k, arg.v, v)
		}
	}
}

func TestSetGet(t *testing.T) {
	m := maps.NewMap()
	for _, arg := range args_set {
		m.Set(arg.k, arg.v)
	}
	for _, arg := range args_expect {
		if v := m.Get(arg.k); v != arg.v {
			t.Fatalf("for key '%s', expect '%v', but was '%v'", arg.k, arg.v, v)
		}
	}
}
