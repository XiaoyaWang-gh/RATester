package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithQuery_1(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want simpleConditionOption
	}{
		{
			name: "TestWithQuery",
			args: args{
				key:   "testKey",
				value: "testValue",
			},
			want: WithQuery("testKey", "testValue"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithQuery(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
