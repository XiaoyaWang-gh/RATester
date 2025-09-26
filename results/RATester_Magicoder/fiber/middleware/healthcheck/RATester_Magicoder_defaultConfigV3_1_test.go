package healthcheck

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestdefaultConfigV3_1(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			name: "Test case 1",
			want: Config{
				Probe: defaultProbe,
			},
		},
		{
			name: "Test case 2",
			want: Config{
				Probe: func(c fiber.Ctx) bool {
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

			if got := defaultConfigV3(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultConfigV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
