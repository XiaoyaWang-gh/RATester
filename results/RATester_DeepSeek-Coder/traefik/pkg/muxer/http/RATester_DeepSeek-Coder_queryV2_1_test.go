package http

import (
	"fmt"
	"testing"
)

func TestQueryV2_1(t *testing.T) {
	type args struct {
		tree  *matchersTree
		query []string
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

			if err := queryV2(tt.args.tree, tt.args.query...); (err != nil) != tt.wantErr {
				t.Errorf("queryV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
