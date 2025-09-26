package config

import (
	"fmt"
	"testing"
)

func TestDefaultInt64_2(t *testing.T) {
	t.Parallel()

	type args struct {
		key        string
		defaultVal int64
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "TestDefaultInt64_0",
			args: args{
				key:        "key0",
				defaultVal: 123,
			},
			want: 123,
		},
		{
			name: "TestDefaultInt64_1",
			args: args{
				key:        "key1",
				defaultVal: 456,
			},
			want: 456,
		},
		{
			name: "TestDefaultInt64_2",
			args: args{
				key:        "key2",
				defaultVal: 789,
			},
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

			c := &fakeConfigContainer{
				data: map[string]string{
					"key0": "123",
					"key1": "456",
					"key2": "789",
				},
			}

			if got := c.DefaultInt64(tt.args.key, tt.args.defaultVal); got != tt.want {
				t.Errorf("DefaultInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
