package codegen

import (
	"fmt"
	"testing"
)

func TestInStr_1(t *testing.T) {
	tests := []struct {
		name string
		m    Method
		want string
	}{
		{
			name: "empty in",
			m:    Method{In: []string{}},
			want: "()",
		},
		{
			name: "single in",
			m:    Method{In: []string{"string"}},
			want: "(arg0 string)",
		},
		{
			name: "multiple in",
			m:    Method{In: []string{"string", "int"}},
			want: "(arg0 string, arg1 int)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.inStr(); got != tt.want {
				t.Errorf("Method.inStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
