package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttCSSStr_1(t *testing.T) {
	tests := []struct {
		name string
		c    context
		s    []byte
		want context
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

			if got, _ := tCSSStr(tt.c, tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tCSSStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
