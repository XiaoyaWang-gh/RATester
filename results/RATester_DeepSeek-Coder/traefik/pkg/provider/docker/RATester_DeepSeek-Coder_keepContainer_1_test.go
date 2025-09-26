package docker

import (
	"context"
	"fmt"
	"testing"
)

func TestKeepContainer_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		container dockerData
	}
	tests := []struct {
		name string
		args args
		want bool
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

			p := &DynConfBuilder{}
			if got := p.keepContainer(tt.args.ctx, tt.args.container); got != tt.want {
				t.Errorf("DynConfBuilder.keepContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
