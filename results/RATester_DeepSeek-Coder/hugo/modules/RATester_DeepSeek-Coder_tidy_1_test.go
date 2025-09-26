package modules

import (
	"fmt"
	"testing"
)

func TestTidy_1(t *testing.T) {
	type args struct {
		mods      Modules
		goModOnly bool
	}
	tests := []struct {
		name    string
		c       *Client
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

			if err := tt.c.tidy(tt.args.mods, tt.args.goModOnly); (err != nil) != tt.wantErr {
				t.Errorf("Client.tidy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
