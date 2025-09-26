package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttJS_1(t *testing.T) {
	tests := []struct {
		name  string
		c     context
		s     []byte
		want  context
		want1 int
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

			got, got1 := tJS(tt.c, tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tJS() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tJS() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
