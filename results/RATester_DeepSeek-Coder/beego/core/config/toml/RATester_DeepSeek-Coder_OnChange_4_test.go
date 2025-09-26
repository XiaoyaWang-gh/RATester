package toml

import (
	"fmt"
	"testing"
)

func TestOnChange_4(t *testing.T) {
	type args struct {
		key string
		fn  func(value string)
	}
	tests := []struct {
		name string
		c    *configContainer
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

			tt.c.OnChange(tt.args.key, tt.args.fn)
		})
	}
}
