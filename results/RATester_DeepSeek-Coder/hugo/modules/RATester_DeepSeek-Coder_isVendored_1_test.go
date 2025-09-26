package modules

import (
	"fmt"
	"testing"
)

func TestIsVendored_1(t *testing.T) {
	type fields struct {
		Client    *Client
		err       error
		skipTidy  bool
		collected *collected
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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

			c := &collector{
				Client:    tt.fields.Client,
				err:       tt.fields.err,
				skipTidy:  tt.fields.skipTidy,
				collected: tt.fields.collected,
			}
			if got := c.isVendored(tt.args.dir); got != tt.want {
				t.Errorf("collector.isVendored() = %v, want %v", got, tt.want)
			}
		})
	}
}
