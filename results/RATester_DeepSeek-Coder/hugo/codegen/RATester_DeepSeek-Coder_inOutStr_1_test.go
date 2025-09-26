package codegen

import (
	"fmt"
	"testing"
)

func TestInOutStr_1(t *testing.T) {
	tests := []struct {
		name string
		m    Method
		want string
	}{
		{
			name: "no args",
			m:    Method{In: []string{}},
			want: "()",
		},
		{
			name: "one arg",
			m:    Method{In: []string{"string"}},
			want: "(arg0)",
		},
		{
			name: "multiple args",
			m:    Method{In: []string{"string", "int", "interface{}"}},
			want: "(arg0, arg1, arg2)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.inOutStr(); got != tt.want {
				t.Errorf("Method.inOutStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
