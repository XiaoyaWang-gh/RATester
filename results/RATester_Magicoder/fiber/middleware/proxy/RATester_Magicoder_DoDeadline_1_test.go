package proxy

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func TestDoDeadline_1(t *testing.T) {
	type args struct {
		c        fiber.Ctx
		addr     string
		deadline time.Time
		clients  []*fasthttp.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			if err := DoDeadline(tt.args.c, tt.args.addr, tt.args.deadline, tt.args.clients...); (err != nil) != tt.wantErr {
				t.Errorf("DoDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
