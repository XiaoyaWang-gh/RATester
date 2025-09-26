package config

import (
	"fmt"
	"testing"
)

func TestDefaultString_3(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal string
	}

	tests := []struct {
		name string
		c    *IniConfigContainer
		args args
		want string
	}{
		{
			name: "TestDefaultString_0",
			c:    &IniConfigContainer{},
			args: args{
				key:        "key",
				defaultVal: "default",
			},
			want: "default",
		},
		{
			name: "TestDefaultString_1",
			c:    &IniConfigContainer{},
			args: args{
				key:        "key",
				defaultVal: "default",
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

			if got := tt.c.DefaultString(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}
