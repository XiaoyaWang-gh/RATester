package yaml

import (
	"fmt"
	"sync"
	"testing"
)

func TestDefaultBool_1(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal bool
	}

	tests := []struct {
		name string
		c    *ConfigContainer
		args args
		want bool
	}{
		{
			name: "TestDefaultBool_0",
			c:    &ConfigContainer{data: map[string]interface{}{"key": true}},
			args: args{key: "key", defaultVal: false},
			want: true,
		},
		{
			name: "TestDefaultBool_1",
			c:    &ConfigContainer{data: map[string]interface{}{"key": false}},
			args: args{key: "key", defaultVal: true},
			want: false,
		},
		{
			name: "TestDefaultBool_2",
			c:    &ConfigContainer{data: map[string]interface{}{"key": "value"}},
			args: args{key: "key", defaultVal: true},
			want: true,
		},
		{
			name: "TestDefaultBool_3",
			c:    &ConfigContainer{data: map[string]interface{}{"key": 1}},
			args: args{key: "key", defaultVal: false},
			want: false,
		},
		{
			name: "TestDefaultBool_4",
			c:    &ConfigContainer{data: map[string]interface{}{"key": 0}, RWMutex: sync.RWMutex{}},
			args: args{key: "key", defaultVal: true},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.DefaultBool(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
