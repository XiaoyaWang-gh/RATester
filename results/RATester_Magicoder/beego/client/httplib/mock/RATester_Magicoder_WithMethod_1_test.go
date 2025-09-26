package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithMethod_1(t *testing.T) {
	type args struct {
		method string
	}
	tests := []struct {
		name string
		args args
		want simpleConditionOption
	}{
		{
			name: "Test WithMethod",
			args: args{
				method: "GET",
			},
			want: func(sc *SimpleCondition) {
				sc.method = "GET"
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

			if got := WithMethod(tt.args.method); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
