package health

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestDoTCPCheck_1(t *testing.T) {
	type fields struct {
		checkType      string
		interval       time.Duration
		timeout        time.Duration
		maxFailedTimes int
		addr           string
		url            string
		header         http.Header
		failedTimes    uint64
		statusOK       bool
		statusNormalFn func()
		statusFailedFn func()
		ctx            context.Context
		cancel         context.CancelFunc
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
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

			monitor := &Monitor{
				checkType:      tt.fields.checkType,
				interval:       tt.fields.interval,
				timeout:        tt.fields.timeout,
				maxFailedTimes: tt.fields.maxFailedTimes,
				addr:           tt.fields.addr,
				url:            tt.fields.url,
				header:         tt.fields.header,
				failedTimes:    tt.fields.failedTimes,
				statusOK:       tt.fields.statusOK,
				statusNormalFn: tt.fields.statusNormalFn,
				statusFailedFn: tt.fields.statusFailedFn,
				ctx:            tt.fields.ctx,
				cancel:         tt.fields.cancel,
			}
			if err := monitor.doTCPCheck(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Monitor.doTCPCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
