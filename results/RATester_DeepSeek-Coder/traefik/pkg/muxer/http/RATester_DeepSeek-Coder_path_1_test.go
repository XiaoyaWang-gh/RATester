package http

import (
	"fmt"
	"testing"
)

func TestPath_1(t *testing.T) {
	type args struct {
		tree *matchersTree
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid path",
			args: args{
				tree: &matchersTree{},
				path: "/valid/path",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid path",
			args: args{
				tree: &matchersTree{},
				path: "invalid/path",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := path(tt.args.tree, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("path() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
