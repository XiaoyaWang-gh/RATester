package healthcheck

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestDefaultProbe_1(t *testing.T) {
	type args struct {
		ctx fiber.Ctx
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := defaultProbe(tt.args.ctx); got != tt.want {
				t.Errorf("defaultProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}
