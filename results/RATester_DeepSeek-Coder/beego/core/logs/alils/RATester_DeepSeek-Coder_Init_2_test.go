package alils

import (
	"fmt"
	"testing"
)

func TestInit_2(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		c       *aliLSWriter
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

			c := &aliLSWriter{}
			if err := c.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("aliLSWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
