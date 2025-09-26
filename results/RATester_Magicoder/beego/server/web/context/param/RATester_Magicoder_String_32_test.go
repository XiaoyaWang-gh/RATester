package param

import (
	"fmt"
	"testing"
)

func TestString_32(t *testing.T) {
	tests := []struct {
		name         string
		in           paramType
		required     bool
		defaultValue string
		want         string
	}{
		{
			name:         "test",
			in:           path,
			required:     true,
			defaultValue: "",
			want:         "param.New(\"test\", param.InPath, param.IsRequired)",
		},
		{
			name:         "test2",
			in:           body,
			required:     false,
			defaultValue: "default",
			want:         "param.New(\"test2\", param.InBody, param.Default(\"default\"))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			mp := &MethodParam{
				name:         tt.name,
				in:           tt.in,
				required:     tt.required,
				defaultValue: tt.defaultValue,
			}
			if got := mp.String(); got != tt.want {
				t.Errorf("MethodParam.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
