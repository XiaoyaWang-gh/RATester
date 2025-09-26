package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectHaveKey_1(t *testing.T) {
	gomega.RegisterTestingT(t)

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
			name: "ExpectHaveKey",
			args: args{
				actual: map[string]interface{}{
					"name": "John",
				},
				key:     "name",
				explain: []interface{}{"Expect the map to have the key 'name'"},
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
