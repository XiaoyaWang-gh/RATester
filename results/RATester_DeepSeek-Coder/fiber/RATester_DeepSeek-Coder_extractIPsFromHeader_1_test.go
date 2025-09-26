package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtractIPsFromHeader_1(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		c    *DefaultCtx
		args args
		want []string
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

			if got := tt.c.extractIPsFromHeader(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultCtx.extractIPsFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
