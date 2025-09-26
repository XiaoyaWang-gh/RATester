package env

import (
	"fmt"
	"testing"
)

func TestSet_32(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				key:   "key1",
				value: "value1",
			},
		},
		{
			name: "Test case 2",
			args: args{
				key:   "key2",
				value: "value2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Set(tt.args.key, tt.args.value)
			val, ok := env.Load(tt.args.key)
			if !ok {
				t.Errorf("Key not found in the map")
			}
			if val != tt.args.value {
				t.Errorf("Expected value %s, got %s", tt.args.value, val)
			}
		})
	}
}
