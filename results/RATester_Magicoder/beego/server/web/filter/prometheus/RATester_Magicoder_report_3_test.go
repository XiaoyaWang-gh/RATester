package prometheus

import (
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/prometheus/client_golang/prometheus"
)

func Testreport_3(t *testing.T) {
	type args struct {
		dur time.Duration
		ctx *context.Context
		vec prometheus.ObserverVec
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

			report(tt.args.dur, tt.args.ctx, tt.args.vec)
		})
	}
}
