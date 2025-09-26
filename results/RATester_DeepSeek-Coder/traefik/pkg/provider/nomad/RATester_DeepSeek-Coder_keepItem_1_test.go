package nomad

import (
	"context"
	"fmt"
	"testing"
)

func TestKeepItem_1(t *testing.T) {
	type args struct {
		ctx context.Context
		i   item
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

			p := &Provider{}
			if got := p.keepItem(tt.args.ctx, tt.args.i); got != tt.want {
				t.Errorf("Provider.keepItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
