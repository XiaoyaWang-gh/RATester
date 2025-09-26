package gin

import (
	"fmt"
	"testing"
)

func TestServeError_1(t *testing.T) {
	type args struct {
		c              *Context
		code           int
		defaultMessage []byte
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

			serveError(tt.args.c, tt.args.code, tt.args.defaultMessage)
		})
	}
}
