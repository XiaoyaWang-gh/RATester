package http

import (
	"fmt"
	"testing"
)

func TestMethodsV2_1(t *testing.T) {
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

			if err := methodsV2(tt.args.tree, tt.args.methods...); (err != nil) != tt.wantErr {
				t.Errorf("methodsV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
