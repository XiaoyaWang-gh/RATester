package gin

import (
	"fmt"
	"testing"
)

func TestTOML_1(t *testing.T) {
	type args struct {
		code int
		obj  any
	}
	tests := []struct {
		name string
		c    *Context
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

			tt.c.TOML(tt.args.code, tt.args.obj)
		})
	}
}
