package xml

import (
	"fmt"
	"testing"
)

func TestDefaultString_8(t *testing.T) {
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
			c: &ConfigContainer{
				data: map[string]interface{}{
					"key1": "value1",
				},
			},
			args: args{
				key:        "key1",
				defaultVal: "default",
			},
			want: "value1",
		},
		{
			name: "TestDefaultString_1",
			c: &ConfigContainer{
				data: map[string]interface{}{
					"key2": "",
				},
			},
			args: args{
				key:        "key2",
				defaultVal: "default",
			},
			want: "default",
		},
		{
			name: "TestDefaultString_2",
			c: &ConfigContainer{
				data: map[string]interface{}{
					"key3": "value3",
				},
			},
			args: args{
				key:        "key4",
				defaultVal: "default",
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

			if got := tt.c.DefaultString(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}
