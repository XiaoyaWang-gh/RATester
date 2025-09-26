package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetDatas_1(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		f    *FormData
		args args
		want *FormData
	}{
		{
			name: "TestSetDatas",
			f:    &FormData{Args: &fasthttp.Args{}},
			args: args{m: map[string]string{"key1": "value1", "key2": "value2"}},
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

			tt.f.SetDatas(tt.args.m)
			for k, v := range tt.args.m {
				if !tt.f.Has(k) {
					t.Errorf("SetDatas() = %v, want %v", tt.f, tt.want)
				}
				if string(tt.f.Peek(k)) != v {
					t.Errorf("SetDatas() = %v, want %v", tt.f, tt.want)
				}
			}
		})
	}
}
