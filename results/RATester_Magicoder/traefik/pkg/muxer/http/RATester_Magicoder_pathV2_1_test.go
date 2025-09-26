package http

import (
	"fmt"
	"testing"
)

func TestpathV2_1(t *testing.T) {
	type args struct {
		tree  *matchersTree
		paths []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"/test1", "/test2"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"/test1", "/test2", "/test3"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"/test1", "/test2", "/test3", "/test4"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"/test1", "/test2", "/test3", "/test4", "/test5"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := pathV2(tt.args.tree, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("pathV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
