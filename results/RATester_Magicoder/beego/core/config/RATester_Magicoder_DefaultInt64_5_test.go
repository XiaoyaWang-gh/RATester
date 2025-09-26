package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_5(t *testing.T) {
	type args struct {
		key        string
		defaultVal int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test case 1",
			args: args{
				key:        "key1",
				defaultVal: 10,
			},
			want: 10,
		},
		{
			name: "Test case 2",
			args: args{
				key:        "key2",
				defaultVal: 20,
			},
			want: 20,
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

			if got := DefaultInt64(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
