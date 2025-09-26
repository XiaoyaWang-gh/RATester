package http

import (
	"fmt"
	"testing"
)

func TestMethod_1(t *testing.T) {
	type args struct {
		tree    *matchersTree
		methods []string
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

			if err := method(tt.args.tree, tt.args.methods...); (err != nil) != tt.wantErr {
				t.Errorf("method() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
