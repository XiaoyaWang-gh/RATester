package rewrite

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestconfigDefault_18(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test Case 1",
			args: args{
				config: []Config{
					{
						Next: func(fiber.Ctx) bool {
							return true
						},
						Rules: map[string]string{
							"/old": "/new",
						},
					},
				},
			},
			want: Config{
				Next: func(fiber.Ctx) bool {
					return true
				},
				Rules: map[string]string{
					"/old": "/new",
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				config: []Config{},
			},
			want: Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := configDefault(tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
