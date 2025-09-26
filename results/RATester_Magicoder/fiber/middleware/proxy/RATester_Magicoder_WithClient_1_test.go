package proxy

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestWithClient_1(t *testing.T) {
	type args struct {
		cli *fasthttp.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			WithClient(tt.args.cli)
		})
	}
}
