package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestsetRaw_1(t *testing.T) {
	type args struct {
		key string
		raw []byte
		exp time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				key: "key1",
				raw: []byte("value1"),
				exp: time.Hour,
			},
		},
		{
			name: "Test Case 2",
			args: args{
				key: "key2",
				raw: []byte("value2"),
				exp: time.Minute,
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &manager{
				memory: &memory.Storage{},
			}
			m.setRaw(tt.args.key, tt.args.raw, tt.args.exp)
			// Add assertions to check if the method did what it was supposed to do
		})
	}
}
