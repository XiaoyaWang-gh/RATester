package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		input *BeegoInput
		key   string
		val   string
		want  *BeegoInput
	}

	tests := []testStruct{
		{
			input: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "value2", "value3"},
			},
			key: "key2",
			val: "newValue2",
			want: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "newValue2", "value3"},
			},
		},
		{
			input: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3"},
				pvalues: []string{"value1", "value2", "value3"},
			},
			key: "key4",
			val: "newValue4",
			want: &BeegoInput{
				pnames:  []string{"key1", "key2", "key3", "key4"},
				pvalues: []string{"value1", "value2", "value3", "newValue4"},
			},
		},
	}

	for _, test := range tests {
		test.input.SetParam(test.key, test.val)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("SetParam(%s, %s) = %v; want %v", test.key, test.val, test.input, test.want)
		}
	}
}
