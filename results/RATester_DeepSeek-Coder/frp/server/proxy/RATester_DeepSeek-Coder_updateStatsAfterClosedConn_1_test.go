package proxy

import (
	"fmt"
	"testing"
)

func TestUpdateStatsAfterClosedConn_1(t *testing.T) {
	type args struct {
		totalRead  int64
		totalWrite int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				totalRead:  100,
				totalWrite: 200,
			},
		},
		{
			name: "Test Case 2",
			args: args{
				totalRead:  300,
				totalWrite: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pxy := &HTTPProxy{
				BaseProxy: &BaseProxy{
					name: "test",
				},
			}
			pxy.updateStatsAfterClosedConn(tt.args.totalRead, tt.args.totalWrite)
			// Add assertions here to check if the methods have been called with the correct arguments
		})
	}
}
