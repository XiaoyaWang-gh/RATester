package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAddData_1(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name string
		f    *FormData
		args args
	}{
		{
			name: "TestAddData",
			f:    &FormData{Args: &fasthttp.Args{}},
			args: args{
				key: "testKey",
				val: "testVal",
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

			tt.f.AddData(tt.args.key, tt.args.val)
			if !tt.f.Has(tt.args.key) {
				t.Errorf("AddData() = %v, want %v", tt.f.Has(tt.args.key), true)
			}
		})
	}
}
