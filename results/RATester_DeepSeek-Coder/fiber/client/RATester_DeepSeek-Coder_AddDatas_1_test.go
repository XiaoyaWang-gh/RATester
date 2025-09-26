package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAddDatas_1(t *testing.T) {
	type args struct {
		m map[string][]string
	}
	tests := []struct {
		name string
		f    *FormData
		args args
		want *FormData
	}{
		{
			name: "TestAddDatas",
			f:    &FormData{Args: &fasthttp.Args{}},
			args: args{m: map[string][]string{"key1": {"val1", "val2"}, "key2": {"val3", "val4"}}},
			want: &FormData{Args: &fasthttp.Args{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.f.AddDatas(tt.args.m)
			if !reflect.DeepEqual(tt.f, tt.want) {
				t.Errorf("AddDatas() = %v, want %v", tt.f, tt.want)
			}
		})
	}
}
