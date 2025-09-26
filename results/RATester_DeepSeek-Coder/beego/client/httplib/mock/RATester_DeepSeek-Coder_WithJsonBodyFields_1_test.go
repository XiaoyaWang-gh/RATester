package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithJsonBodyFields_1(t *testing.T) {
	type args struct {
		field string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want simpleConditionOption
	}{
		{
			name: "TestWithJsonBodyFields_String",
			args: args{
				field: "testField",
				value: "testValue",
			},
			want: WithJsonBodyFields("testField", "testValue"),
		},
		{
			name: "TestWithJsonBodyFields_Int",
			args: args{
				field: "testField",
				value: 123,
			},
			want: WithJsonBodyFields("testField", 123),
		},
		{
			name: "TestWithJsonBodyFields_Bool",
			args: args{
				field: "testField",
				value: true,
			},
			want: WithJsonBodyFields("testField", true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithJsonBodyFields(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithJsonBodyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
