package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestFlowAddHost_1(t *testing.T) {
	type args struct {
		host *file.Host
		in   int64
		out  int64
	}

	testHost := &file.Host{
		Flow: &file.Flow{
			ExportFlow: 0,
			InletFlow:  0,
		},
	}

	tests := []struct {
		name string
		s    *BaseServer
		args args
	}{
		{
			name: "TestFlowAddHost",
			s:    &BaseServer{},
			args: args{
				host: testHost,
				in:   10,
				out:  20,
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

			tt.s.FlowAddHost(tt.args.host, tt.args.in, tt.args.out)
			if tt.args.host.Flow.ExportFlow != tt.args.out {
				t.Errorf("FlowAddHost() = %v, want %v", tt.args.host.Flow.ExportFlow, tt.args.out)
			}
			if tt.args.host.Flow.InletFlow != tt.args.in {
				t.Errorf("FlowAddHost() = %v, want %v", tt.args.host.Flow.InletFlow, tt.args.in)
			}
		})
	}
}
