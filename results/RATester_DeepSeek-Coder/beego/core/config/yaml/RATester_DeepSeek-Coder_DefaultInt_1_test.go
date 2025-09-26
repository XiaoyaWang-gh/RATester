package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultInt_1(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal int
	}

	tests := []struct {
		name string
		c    *ConfigContainer
		args args
		want int
	}{
		{
			name: "TestDefaultInt_1",
			c:    &ConfigContainer{data: map[string]interface{}{"key1": 123}},
			args: args{key: "key1", defaultVal: 456},
			want: 123,
		},
		{
			name: "TestDefaultInt_2",
			c:    &ConfigContainer{data: map[string]interface{}{"key2": "abc"}},
			args: args{key: "key2", defaultVal: 789},
			want: 789,
		},
		{
			name: "TestDefaultInt_3",
			c:    &ConfigContainer{data: map[string]interface{}{"key3": nil}},
			args: args{key: "key3", defaultVal: 0},
			want: 0,
		},
		{
			name: "TestDefaultInt_4",
			c:    &ConfigContainer{data: map[string]interface{}{}},
			args: args{key: "key4", defaultVal: 111},
			want: 111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.DefaultInt(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
