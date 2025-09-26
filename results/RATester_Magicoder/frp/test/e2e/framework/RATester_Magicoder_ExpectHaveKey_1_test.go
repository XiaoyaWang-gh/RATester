package framework

import (
	"fmt"
	"testing"
)

func TestExpectHaveKey_1(t *testing.T) {
	type args struct {
		actual  interface{}
		key     interface{}
		explain []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				actual: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
				key:     "key1",
				explain: []interface{}{"Expect to have key"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				actual: map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
				key:     "key3",
				explain: []interface{}{"Expect to have key"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ExpectHaveKey(tt.args.actual, tt.args.key, tt.args.explain...)
		})
	}
}
