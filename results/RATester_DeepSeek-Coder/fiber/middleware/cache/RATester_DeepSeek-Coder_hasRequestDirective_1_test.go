package cache

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestHasRequestDirective_1(t *testing.T) {
	type args struct {
		c         fiber.Ctx
		directive string
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

			if got := hasRequestDirective(tt.args.c, tt.args.directive); got != tt.want {
				t.Errorf("hasRequestDirective() = %v, want %v", got, tt.want)
			}
		})
	}
}
