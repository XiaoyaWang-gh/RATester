package proxy

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func TestDoRedirects_1(t *testing.T) {
	type args struct {
		c                 fiber.Ctx
		addr              string
		maxRedirectsCount int
		clients           []*fasthttp.Client
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

			if err := DoRedirects(tt.args.c, tt.args.addr, tt.args.maxRedirectsCount, tt.args.clients...); (err != nil) != tt.wantErr {
				t.Errorf("DoRedirects() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
