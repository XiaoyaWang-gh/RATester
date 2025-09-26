package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockMethod_1(t *testing.T) {
	type args struct {
		method string
		resp   []interface{}
	}
	tests := []struct {
		name string
		args args
		want *Mock
	}{
		{
			name: "Test case 1",
			args: args{
				method: "GET",
				resp:   []interface{}{"OK"},
			},
			want: &Mock{
				cond: &SimpleCondition{
					tableName: "",
					method:    "GET",
				},
				resp: []interface{}{"OK"},
				cb:   nil,
			},
		},
		{
			name: "Test case 2",
			args: args{
				method: "POST",
				resp:   []interface{}{"Created"},
			},
			want: &Mock{
				cond: &SimpleCondition{
					tableName: "",
					method:    "POST",
				},
				resp: []interface{}{"Created"},
				cb:   nil,
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

			if got := MockMethod(tt.args.method, tt.args.resp...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
