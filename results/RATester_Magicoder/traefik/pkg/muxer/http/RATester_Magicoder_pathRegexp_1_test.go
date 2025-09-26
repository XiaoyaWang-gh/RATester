package http

import (
	"fmt"
	"testing"
)

func TestpathRegexp_1(t *testing.T) {
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
				paths: []string{"/test/path"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"[a-z]+"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"*"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"[a-z]"},
			},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"[a-z]+"},
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

			if err := pathRegexp(tt.args.tree, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("pathRegexp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
