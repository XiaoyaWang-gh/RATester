package mem

import (
	"fmt"
	"testing"
)

func TestNewProxy_6(t *testing.T) {
	type args struct {
		name      string
		proxyType string
	}
	tests := []struct {
		name string
		args args
		want *serverMetrics
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

			m := &serverMetrics{}
			m.NewProxy(tt.args.name, tt.args.proxyType)
		})
	}
}
