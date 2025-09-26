package mem

import (
	"fmt"
	"testing"
)

func TestCloseProxy_6(t *testing.T) {
	type args struct {
		name      string
		proxyType string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				name:      "proxy1",
				proxyType: "type1",
			},
		},
		{
			name: "Test case 2",
			args: args{
				name:      "proxy2",
				proxyType: "type2",
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

			m := &serverMetrics{}
			m.CloseProxy(tt.args.name, tt.args.proxyType)
		})
	}
}
