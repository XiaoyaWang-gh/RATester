package http

import (
	"fmt"
	"testing"
)

func TestHeaderRegexp_1(t *testing.T) {
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

			if err := headerRegexp(tt.args.tree, tt.args.headers...); (err != nil) != tt.wantErr {
				t.Errorf("headerRegexp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
