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
			name: "empty",
			m:    Method{},
			want: "",
		},
		{
			name: "one",
			m:    Method{Out: []string{"string"}},
			want: "string",
		},
		{
			name: "two",
			m:    Method{Out: []string{"string", "int"}},
			want: "(string, int)",
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
