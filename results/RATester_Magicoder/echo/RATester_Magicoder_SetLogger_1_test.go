package echo

import (
	"fmt"
	"testing"
)

func TestSetLogger_1(t *testing.T) {
	type args struct {
		l Logger
	}
	tests := []struct {
		name string
		c    *context
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

			tt.c.SetLogger(tt.args.l)
		})
	}
}
