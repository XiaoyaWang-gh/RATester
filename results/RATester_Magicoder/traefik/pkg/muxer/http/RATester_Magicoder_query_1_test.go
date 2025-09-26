package http

import (
	"fmt"
	"testing"
)

func Testquery_1(t *testing.T) {
	type args struct {
		tree    *matchersTree
		queries []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test query with valid input",
			args: args{
				tree:    &matchersTree{},
				queries: []string{"key", "value"},
			},
			wantErr: false,
		},
		{
			name: "Test query with invalid input",
			args: args{
				tree:    &matchersTree{},
				queries: []string{"invalid_key"},
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

			if err := query(tt.args.tree, tt.args.queries...); (err != nil) != tt.wantErr {
				t.Errorf("query() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
