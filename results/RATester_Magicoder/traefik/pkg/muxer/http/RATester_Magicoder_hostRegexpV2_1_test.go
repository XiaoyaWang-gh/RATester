package http

import (
	"fmt"
	"testing"
)

func TesthostRegexpV2_1(t *testing.T) {
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
				hosts: []string{"example.com", "test.com"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"example.com", "test.com", "invalid host"},
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"example.com", "test.com", "invalid host"},
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

			if err := hostRegexpV2(tt.args.tree, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("hostRegexpV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
