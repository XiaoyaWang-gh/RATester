package csrf

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3/internal/memory"
)

func TestSetRaw_2(t *testing.T) {
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
				key: "testKey",
				raw: []byte("testValue"),
				exp: time.Duration(10),
			},
		},
		{
			name: "Test Case 2",
			args: args{
				key: "testKey2",
				raw: []byte("testValue2"),
				exp: time.Duration(20),
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

			m := &storageManager{
				pool:   sync.Pool{},
				memory: &memory.Storage{},
			}
			m.setRaw(tt.args.key, tt.args.raw, tt.args.exp)
		})
	}
}
