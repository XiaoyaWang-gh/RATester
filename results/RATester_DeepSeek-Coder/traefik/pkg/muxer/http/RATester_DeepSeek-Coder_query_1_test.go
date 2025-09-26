package http

import (
	"fmt"
	"testing"
)

func TestQuery_1(t *testing.T) {
	type args struct {
		tree    *matchersTree
		queries []string
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

			if err := query(tt.args.tree, tt.args.queries...); (err != nil) != tt.wantErr {
				t.Errorf("query() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
