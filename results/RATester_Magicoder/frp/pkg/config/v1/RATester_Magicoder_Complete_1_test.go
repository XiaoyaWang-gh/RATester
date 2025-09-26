package v1

import (
	"fmt"
	"testing"
)

func TestComplete_1(t *testing.T) {
	type args struct {
		g *ClientCommonConfig
	}
	tests := []struct {
		name string
		c    *XTCPVisitorConfig
		args args
		want *XTCPVisitorConfig
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

			tt.c.Complete(tt.args.g)
		})
	}
}
