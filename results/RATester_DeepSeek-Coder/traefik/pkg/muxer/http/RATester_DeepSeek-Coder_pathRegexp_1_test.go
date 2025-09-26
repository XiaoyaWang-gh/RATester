package http

import (
	"fmt"
	"testing"
)

func TestPathRegexp_1(t *testing.T) {
	type args struct {
		tree  *matchersTree
		paths []string
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

			if err := pathRegexp(tt.args.tree, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("pathRegexp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
