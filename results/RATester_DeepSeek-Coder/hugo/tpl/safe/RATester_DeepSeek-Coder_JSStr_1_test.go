package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestJSStr_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    template.JSStr
		wantErr bool
	}{
		{
			name:  "string input",
			input: "test",
			want:  template.JSStr("test"),
		},
		{
			name:  "int input",
			input: 123,
			want:  template.JSStr("123"),
		},
		{
			name:  "float input",
			input: 123.456,
			want:  template.JSStr("123.456"),
		},
		{
			name:  "bool input",
			input: true,
			want:  template.JSStr("true"),
		},
		{
			name:  "nil input",
			input: nil,
			want:  template.JSStr("<nil>"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.JSStr(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
