package maps

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		c    *Scratch
		args args
		want string
	}{
		{
			name: "Set a value",
			c:    &Scratch{values: make(map[string]any)},
			args: args{
				key:   "key1",
				value: "value1",
			},
			want: "",
		},
		{
			name: "Set another value",
			c:    &Scratch{values: make(map[string]any)},
			args: args{
				key:   "key2",
				value: 123,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.Set(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("Scratch.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
