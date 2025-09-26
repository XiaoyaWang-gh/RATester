package proxy

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func TestDo_1(t *testing.T) {
	type args struct {
		c       fiber.Ctx
		addr    string
		clients []*fasthttp.Client
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

			if err := Do(tt.args.c, tt.args.addr, tt.args.clients...); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
