package config

import (
	"fmt"
	"testing"
)

func TestDefaultString_2(t *testing.T) {
	type args struct {
		key        string
		defaultVal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				key:        "key1",
				defaultVal: "default1",
			},
			want: "value1",
		},
		{
			name: "Test case 2",
			args: args{
				key:        "key2",
				defaultVal: "default2",
			},
			want: "default2",
		},
		{
			name: "Test case 3",
			args: args{
				key:        "key3",
				defaultVal: "default3",
			},
			want: "default3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DefaultString(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}
