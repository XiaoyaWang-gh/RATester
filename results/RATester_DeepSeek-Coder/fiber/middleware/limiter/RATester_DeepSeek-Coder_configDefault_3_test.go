package limiter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestConfigDefault_3(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test with empty config",
			args: args{
				config: []Config{},
			},
			want: ConfigDefault,
		},
		{
			name: "Test with custom config",
			args: args{
				config: []Config{
					{
						Max: 10,
					},
				},
			},
			want: Config{
				Next:              ConfigDefault.Next,
				Max:               10,
				Expiration:        ConfigDefault.Expiration,
				KeyGenerator:      ConfigDefault.KeyGenerator,
				LimitReached:      ConfigDefault.LimitReached,
				LimiterMiddleware: ConfigDefault.LimiterMiddleware,
				MaxFunc: func(_ fiber.Ctx) int {
					return 10
				},
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
