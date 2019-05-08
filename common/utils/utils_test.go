package utils_test

import (
	"rest-demo/common/utils"
	"testing"
)

type Test struct {
	in  interface{}
	out interface{}
}

type condition struct {
	name   string
	fields []string
}

// TestContains check if Contains works well
func Test_Contains(t *testing.T) {
	var tests = []Test{
		{condition{
			name:   "a",
			fields: []string{"a", "b", "c"},
		}, true},
		{condition{
			name:   "a",
			fields: []string{"b", "b", "d"},
		}, false},
	}
	for i, test := range tests {
		b := utils.Contains(test.in.(condition).name, test.in.(condition).fields)
		if b != test.out {
			t.Errorf("#%d: Contains(%s,%v)=%v; want %v", i, test.in.(condition).name, test.in.(condition).fields, b, test.out)
		}
	}
}
