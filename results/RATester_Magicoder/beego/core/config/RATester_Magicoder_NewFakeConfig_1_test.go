package config

import (
	"fmt"
	"testing"
)

func TestNewFakeConfig_1(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				key:   "test_key",
				value: "test_value",
			},
			want: "test_value",
		},
		{
			name: "Test case 2",
			args: args{
				key:   "test_key_2",
				value: "test_value_2",
			},
			want: "test_value_2",
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

			config := NewFakeConfig()
			config.Set(tt.args.key, tt.args.value)
			got, _ := config.String(tt.args.key)
			if got != tt.want {
				t.Errorf("NewFakeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
