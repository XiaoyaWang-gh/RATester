package metrics

import (
	"fmt"
	"testing"
)

func TestAddTrafficOut_2(t *testing.T) {
	type args struct {
		clientID   string
		connection string
		bytes      int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestAddTrafficOut",
			args: args{
				clientID:   "client1",
				connection: "conn1",
				bytes:      1024,
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

			noopServerMetrics{}.AddTrafficOut(tt.args.clientID, tt.args.connection, tt.args.bytes)
		})
	}
}
