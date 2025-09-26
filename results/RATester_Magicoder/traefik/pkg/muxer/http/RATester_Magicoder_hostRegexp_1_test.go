package http

import (
	"fmt"
	"testing"
)

func TesthostRegexp_1(t *testing.T) {
	type args struct {
		tree  *matchersTree
		hosts []string
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
				hosts: []string{"example.com"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"example.com", "invalid"},
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"example.com", "invalid", "another.com"},
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

			if err := hostRegexp(tt.args.tree, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("hostRegexp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
