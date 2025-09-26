package http

import (
	"fmt"
	"testing"
)

func TestHeader_7(t *testing.T) {
	type args struct {
		tree    *matchersTree
		headers []string
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

			if err := header(tt.args.tree, tt.args.headers...); (err != nil) != tt.wantErr {
				t.Errorf("header() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
