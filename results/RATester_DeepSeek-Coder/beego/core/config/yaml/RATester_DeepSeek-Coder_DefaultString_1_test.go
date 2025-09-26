package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultString_1(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal string
	}

	tests := []struct {
		name string
		c    *ConfigContainer
		args args
		want string
	}{
		{
			name: "TestDefaultString_0",
			c:    &ConfigContainer{data: map[string]interface{}{"key": "value"}},
			args: args{key: "key", defaultVal: "default"},
			want: "value",
		},
		{
			name: "TestDefaultString_1",
			c:    &ConfigContainer{data: map[string]interface{}{}},
			args: args{key: "key", defaultVal: "default"},
			want: "default",
		},
		{
			name: "TestDefaultString_2",
			c:    &ConfigContainer{data: map[string]interface{}{"key": ""}},
			args: args{key: "key", defaultVal: "default"},
			want: "default",
		},
		{
			name: "TestDefaultString_3",
			c:    &ConfigContainer{data: map[string]interface{}{"key": nil}},
			args: args{key: "key", defaultVal: "default"},
			want: "default",
		},
		{
			name: "TestDefaultString_4",
			c:    &ConfigContainer{data: map[string]interface{}{"key": "0"}},
			args: args{key: "key", defaultVal: "default"},
			want: "0",
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
