package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestDelDatas_1(t *testing.T) {
	type args struct {
		key []string
	}
	tests := []struct {
		name string
		f    *FormData
		args args
	}{
		{
			name: "TestDelDatas",
			f:    &FormData{Args: &fasthttp.Args{}},
			args: args{key: []string{"key1", "key2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.f.DelDatas(tt.args.key...)
			for _, v := range tt.args.key {
				if tt.f.Has(v) {
					t.Errorf("DelDatas() = %v, want %v", tt.f.Has(v), false)
				}
			}
		})
	}
}
