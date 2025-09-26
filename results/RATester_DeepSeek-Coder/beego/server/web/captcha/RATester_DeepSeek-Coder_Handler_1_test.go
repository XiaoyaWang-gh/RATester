package captcha

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestHandler_1(t *testing.T) {
	type args struct {
		ctx *context.Context
	}
	tests := []struct {
		name string
		args args
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

			c := &Captcha{}
			c.Handler(tt.args.ctx)
		})
	}
}
