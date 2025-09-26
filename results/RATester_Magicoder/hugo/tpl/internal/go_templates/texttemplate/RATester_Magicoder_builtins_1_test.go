package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testbuiltins_1(t *testing.T) {
	tests := []struct {
		name string
		want FuncMap
	}{
		{
			name: "Test builtins function",
			want: builtins(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := builtins(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builtins() = %v, want %v", got, tt.want)
			}
		})
	}
}
