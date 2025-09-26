package http

import (
	"fmt"
	"testing"
)

func Testpath_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"/test"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				tree:  &matchersTree{},
				paths: []string{"test"},
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := path(tt.args.tree, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("path() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
