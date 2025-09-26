package web

import (
	"fmt"
	"testing"
	"time"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestLogAccess_1(t *testing.T) {
	type args struct {
		ctx        *beecontext.Context
		startTime  *time.Time
		statusCode int
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

			LogAccess(tt.args.ctx, tt.args.startTime, tt.args.statusCode)
		})
	}
}
