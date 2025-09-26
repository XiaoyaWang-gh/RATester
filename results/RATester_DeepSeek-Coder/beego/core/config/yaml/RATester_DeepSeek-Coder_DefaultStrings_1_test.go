package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_1(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal []string
	}

	tests := []struct {
		name string
		c    *ConfigContainer
		args args
		want []string
	}{
		{
			name: "TestDefaultStrings_0",
			c:    &ConfigContainer{data: map[string]interface{}{"key": []string{"value1", "value2"}}},
			args: args{key: "key", defaultVal: []string{"default1", "default2"}},
			want: []string{"value1", "value2"},
		},
		{
			name: "TestDefaultStrings_1",
			c:    &ConfigContainer{data: map[string]interface{}{"key": nil}},
			args: args{key: "key", defaultVal: []string{"default1", "default2"}},
			want: []string{"default1", "default2"},
		},
		{
			name: "TestDefaultStrings_2",
			c:    &ConfigContainer{data: map[string]interface{}{"key": "value"}},
			args: args{key: "key", defaultVal: []string{"default1", "default2"}},
			want: []string{"default1", "default2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.DefaultStrings(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
