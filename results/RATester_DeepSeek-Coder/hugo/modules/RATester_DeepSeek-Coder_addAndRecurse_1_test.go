package modules

import (
	"fmt"
	"testing"
)

func TestAddAndRecurse_1(t *testing.T) {
	type fields struct {
		Client    *Client
		err       error
		skipTidy  bool
		collected *collected
	}
	type args struct {
		owner *moduleAdapter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
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
			if err := c.addAndRecurse(tt.args.owner); (err != nil) != tt.wantErr {
				t.Errorf("collector.addAndRecurse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
