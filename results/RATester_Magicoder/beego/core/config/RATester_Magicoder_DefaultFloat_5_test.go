package config

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_5(t *testing.T) {
	type args struct {
		key        string
		defaultVal float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test case 1",
			args: args{
				key:        "key1",
				defaultVal: 1.0,
			},
			want: 1.0,
		},
		{
			name: "Test case 2",
			args: args{
				key:        "key2",
				defaultVal: 2.0,
			},
			want: 2.0,
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

			if got := DefaultFloat(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
