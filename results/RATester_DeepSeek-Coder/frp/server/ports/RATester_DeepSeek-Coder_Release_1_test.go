package ports

import (
	"fmt"
	"testing"
)

func TestRelease_1(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{port: 8080},
		},
		{
			name: "Test case 2",
			args: args{port: 8081},
		},
		{
			name: "Test case 3",
			args: args{port: 8082},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			pm := &Manager{
				reservedPorts: make(map[string]*PortCtx),
				usedPorts:     make(map[int]*PortCtx),
				freePorts:     make(map[int]struct{}),
				bindAddr:      "localhost",
				netType:       "tcp",
			}
			pm.Release(tt.args.port)
		})
	}
}
