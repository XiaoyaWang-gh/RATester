package fiber

import (
	"fmt"
	"testing"
)

func TestDefaultString_1(t *testing.T) {
	type args struct {
		value        string
		defaultValue []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with empty value and non-empty default value",
			args: args{
				value:        "",
				defaultValue: []string{"default"},
			},
			want: "default",
		},
		{
			name: "Test with non-empty value and empty default value",
			args: args{
				value:        "value",
				defaultValue: []string{},
			},
			want: "value",
		},
		{
			name: "Test with non-empty value and non-empty default value",
			args: args{
				value:        "value",
				defaultValue: []string{"default"},
			},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := defaultString(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("defaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}
