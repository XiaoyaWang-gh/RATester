package csrf

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestIsFromCookie_1(t *testing.T) {
	type args struct {
		extractor any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				extractor: FromCookie,
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				extractor: func(c fiber.Ctx) (string, error) {
					return "", nil
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isFromCookie(tt.args.extractor); got != tt.want {
				t.Errorf("isFromCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
