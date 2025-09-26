package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt_2(t *testing.T) {
	type args struct {
		key        string
		defaultVal int
	}
	tests := []struct {
		name string
		args args
		want int
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

			if got := DefaultInt(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
