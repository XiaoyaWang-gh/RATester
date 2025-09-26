package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAddParams_1(t *testing.T) {
	type args struct {
		r map[string][]string
	}
	tests := []struct {
		name string
		p    *QueryParam
		args args
		want *QueryParam
	}{
		{
			name: "TestAddParams",
			p:    &QueryParam{Args: &fasthttp.Args{}},
			args: args{r: map[string][]string{"key1": {"value1", "value2"}, "key2": {"value3", "value4"}}},
			want: &QueryParam{Args: &fasthttp.Args{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.p.AddParams(tt.args.r)
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("AddParams() = %v, want %v", tt.p, tt.want)
			}
		})
	}
}
