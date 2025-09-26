package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewQueryM2MerCondition_1(t *testing.T) {
	type args struct {
		tableName string
		name      string
	}
	tests := []struct {
		name string
		args args
		want *QueryM2MerCondition
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

			if got := NewQueryM2MerCondition(tt.args.tableName, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueryM2MerCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}
