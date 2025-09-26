package json

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_7(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal int64
	}

	tests := []struct {
		name string
		c    *JSONConfigContainer
		args args
		want int64
	}{
		{
			name: "TestDefaultInt64_1",
			c:    &JSONConfigContainer{data: map[string]interface{}{"key1": int64(123)}},
			args: args{key: "key1", defaultVal: 456},
			want: 123,
		},
		{
			name: "TestDefaultInt64_2",
			c:    &JSONConfigContainer{data: map[string]interface{}{"key2": "notInt64"}},
			args: args{key: "key2", defaultVal: 456},
			want: 456,
		},
		{
			name: "TestDefaultInt64_3",
			c:    &JSONConfigContainer{data: map[string]interface{}{}},
			args: args{key: "key3", defaultVal: 789},
			want: 789,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.DefaultInt64(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
