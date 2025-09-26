package http

import (
	"fmt"
	"testing"
)

func Testhost_1(t *testing.T) {
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
				hosts: []string{"example.com"},
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

			if err := host(tt.args.tree, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("host() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
