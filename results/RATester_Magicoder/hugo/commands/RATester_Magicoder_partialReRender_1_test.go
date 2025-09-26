package commands

import (
	"fmt"
	"testing"
)

func TestpartialReRender_1(t *testing.T) {
	type args struct {
		urls []string
	}
	tests := []struct {
		name    string
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

			c := &serverCommand{}
			if err := c.partialReRender(tt.args.urls...); (err != nil) != tt.wantErr {
				t.Errorf("partialReRender() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
