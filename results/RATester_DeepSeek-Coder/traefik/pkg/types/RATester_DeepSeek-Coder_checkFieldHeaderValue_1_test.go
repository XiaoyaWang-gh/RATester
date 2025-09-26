package types

import (
	"fmt"
	"testing"
)

func TestCheckFieldHeaderValue_1(t *testing.T) {
	type args struct {
		value        string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				value:        "keep",
				defaultValue: "default",
			},
			want: "keep",
		},
		{
			name: "Test case 2",
			args: args{
				value:        "drop",
				defaultValue: "default",
			},
			want: "drop",
		},
		{
			name: "Test case 3",
			args: args{
				value:        "redact",
				defaultValue: "default",
			},
			want: "redact",
		},
		{
			name: "Test case 4",
			args: args{
				value:        "invalid",
				defaultValue: "default",
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := checkFieldHeaderValue(tt.args.value, tt.args.defaultValue); got != tt.want {
				t.Errorf("checkFieldHeaderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
