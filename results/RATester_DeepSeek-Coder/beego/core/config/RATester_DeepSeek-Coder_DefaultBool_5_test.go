package config

import (
	"fmt"
	"testing"
)

func TestDefaultBool_5(t *testing.T) {
	type args struct {
		key        string
		defaultVal bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				key:        "key1",
				defaultVal: true,
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				key:        "key2",
				defaultVal: false,
			},
			want: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DefaultBool(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
