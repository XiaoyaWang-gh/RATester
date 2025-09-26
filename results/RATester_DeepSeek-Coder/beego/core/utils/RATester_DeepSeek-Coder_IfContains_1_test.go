package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIfContains_1(t *testing.T) {
	type args struct {
		key     interface{}
		action  func(value interface{})
		kvs     *SimpleKVs
		wantKVs *SimpleKVs
	}
	tests := []struct {
		name string
		args args
		want *SimpleKVs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.kvs.IfContains(tt.args.key, tt.args.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleKVs.IfContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
