package cache

import (
	"fmt"
	"testing"
)

func TestDel_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    *manager
		args args
	}{
		{
			name: "Test case 1",
			m:    &manager{},
			args: args{
				key: "test_key",
			},
		},
		{
			name: "Test case 2",
			m:    &manager{},
			args: args{
				key: "test_key_2",
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

			tt.m.del(tt.args.key)
		})
	}
}
