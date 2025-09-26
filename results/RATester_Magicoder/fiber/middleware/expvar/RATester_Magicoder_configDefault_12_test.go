package expvar

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestconfigDefault_12(t *testing.T) {

	tests := []struct {
		name string
		args []Config
		want Config
	}{
		{
			name: "Test with no args",
			args: []Config{},
			want: ConfigDefault,
		},
		{
			name: "Test with args",
			args: []Config{
				{
					Next: func(c fiber.Ctx) bool {
						return true
					},
				},
			},
			want: Config{
				Next: func(c fiber.Ctx) bool {
					return true
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

			if got := configDefault(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
