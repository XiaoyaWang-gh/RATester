package idempotency

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/internal/storage/memory"
)

func TestConfigDefault_8(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test with no config",
			args: args{
				config: []Config{},
			},
			want: Config{
				Lock:      NewMemoryLock(),
				Storage:   memory.New(memory.Config{GCInterval: ConfigDefault.Lifetime / 2}),
				Next:      ConfigDefault.Next,
				KeyHeader: ConfigDefault.KeyHeader,
				Lifetime:  ConfigDefault.Lifetime,
			},
		},
		{
			name: "Test with config",
			args: args{
				config: []Config{
					{
						Lock:      NewMemoryLock(),
						Storage:   memory.New(memory.Config{GCInterval: 10 * time.Minute}),
						Next:      func(c fiber.Ctx) bool { return true },
						KeyHeader: "X-Test-Key",
						Lifetime:  20 * time.Minute,
					},
				},
			},
			want: Config{
				Lock:      NewMemoryLock(),
				Storage:   memory.New(memory.Config{GCInterval: 10 * time.Minute}),
				Next:      func(c fiber.Ctx) bool { return true },
				KeyHeader: "X-Test-Key",
				Lifetime:  20 * time.Minute,
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

			got := configDefault(tt.args.config...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
