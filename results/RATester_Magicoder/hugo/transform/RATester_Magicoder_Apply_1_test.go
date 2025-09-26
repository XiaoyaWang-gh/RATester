package transform

import (
	"fmt"
	"io"
	"testing"
)

func TestApply_1(t *testing.T) {
	type args struct {
		to   io.Writer
		from io.Reader
	}
	tests := []struct {
		name    string
		c       *Chain
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

			if err := tt.c.Apply(tt.args.to, tt.args.from); (err != nil) != tt.wantErr {
				t.Errorf("Chain.Apply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
