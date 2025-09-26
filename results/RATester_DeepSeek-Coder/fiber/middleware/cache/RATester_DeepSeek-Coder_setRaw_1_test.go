package cache

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestSetRaw_1(t *testing.T) {
	t.Parallel()

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
				key: "test_key",
				raw: []byte("test_value"),
				exp: time.Hour,
			},
		},
		{
			name: "Test Case 2",
			args: args{
				key: "test_key_2",
				raw: []byte("test_value_2"),
				exp: time.Minute,
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

			m := &manager{
				pool:   sync.Pool{},
				memory: &memory.Storage{},
			}

			m.setRaw(tt.args.key, tt.args.raw, tt.args.exp)

			got := m.getRaw(tt.args.key)
			if !bytes.Equal(got, tt.args.raw) {
				t.Errorf("Expected %v, got %v", tt.args.raw, got)
			}
		})
	}
}
