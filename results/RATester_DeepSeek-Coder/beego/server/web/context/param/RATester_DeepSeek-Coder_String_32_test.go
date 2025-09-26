package param

import (
	"fmt"
	"testing"
)

func TestString_32(t *testing.T) {
	type fields struct {
		name         string
		in           paramType
		required     bool
		defaultValue string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				name:         "test",
				in:           0,
				required:     true,
				defaultValue: "default",
			},
			want: "param.New(\"test\", param.IsRequired, param.InPath, param.Default(\"default\"))",
		},
		{
			name: "Test case 2",
			fields: fields{
				name:         "test",
				in:           1,
				required:     false,
				defaultValue: "",
			},
			want: "param.New(\"test\", param.InBody)",
		},
		{
			name: "Test case 3",
			fields: fields{
				name:         "test",
				in:           2,
				required:     true,
				defaultValue: "",
			},
			want: "param.New(\"test\", param.IsRequired, param.InHeader)",
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
				name:         tt.fields.name,
				in:           tt.fields.in,
				required:     tt.fields.required,
				defaultValue: tt.fields.defaultValue,
			}
			if got := mp.String(); got != tt.want {
				t.Errorf("MethodParam.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
