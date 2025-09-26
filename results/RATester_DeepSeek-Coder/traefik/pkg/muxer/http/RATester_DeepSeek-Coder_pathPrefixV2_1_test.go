package http

import (
	"fmt"
	"testing"
)

func TestPathPrefixV2_1(t *testing.T) {
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

			if err := pathPrefixV2(tt.args.tree, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("pathPrefixV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
