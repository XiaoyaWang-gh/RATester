package fiber

import (
	"fmt"
	"testing"
)

func TestextractIPFromHeader_1(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		c    *DefaultCtx
		args args
		want string
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

			if got := tt.c.extractIPFromHeader(tt.args.header); got != tt.want {
				t.Errorf("extractIPFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
