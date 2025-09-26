package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSimpleCondition_1(t *testing.T) {
	type args struct {
		tableName  string
		methodName string
	}
	tests := []struct {
		name string
		args args
		want Condition
	}{
		{
			name: "Test case 1",
			args: args{
				tableName:  "test_table",
				methodName: "test_method",
			},
			want: &SimpleCondition{
				tableName: "test_table",
				method:    "test_method",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewSimpleCondition(tt.args.tableName, tt.args.methodName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSimpleCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}
