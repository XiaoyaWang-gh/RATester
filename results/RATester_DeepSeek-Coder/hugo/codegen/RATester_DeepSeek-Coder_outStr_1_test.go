package codegen

import (
	"fmt"
	"testing"
)

func TestOutStr_1(t *testing.T) {
	tests := []struct {
		name string
		m    Method
		want string
	}{
		{
			name: "empty out",
			m:    Method{Out: []string{}},
			want: "",
		},
		{
			name: "single out",
			m:    Method{Out: []string{"error"}},
			want: "error",
		},
		{
			name: "multiple out",
			m:    Method{Out: []string{"error", "int"}},
			want: "(error, int)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.outStr(); got != tt.want {
				t.Errorf("Method.outStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
