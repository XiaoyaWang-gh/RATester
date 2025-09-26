package proxy

import (
	"fmt"
	"testing"
)

func TestUpdateStatsAfterClosedConn_1(t *testing.T) {
	type fields struct {
		BaseProxy *BaseProxy
	}
	type args struct {
		totalRead  int64
		totalWrite int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			pxy := &HTTPProxy{
				BaseProxy: tt.fields.BaseProxy,
			}
			pxy.updateStatsAfterClosedConn(tt.args.totalRead, tt.args.totalWrite)
		})
	}
}
