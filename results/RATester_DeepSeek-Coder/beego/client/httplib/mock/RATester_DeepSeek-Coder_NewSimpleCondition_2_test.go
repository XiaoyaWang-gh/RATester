package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSimpleCondition_2(t *testing.T) {
	type args struct {
		path string
		opts []simpleConditionOption
	}
	tests := []struct {
		name string
		args args
		want *SimpleCondition
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

			if got := NewSimpleCondition(tt.args.path, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimpleCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}
