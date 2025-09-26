package tcp

import (
	"fmt"
	"testing"
)

func TestHostSNIRegexpV2_1(t *testing.T) {
	type args struct {
		tree      *matchersTree
		templates []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				tree:      &matchersTree{},
				templates: []string{"example.com"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree:      &matchersTree{},
				templates: []string{"example.com", "example.org"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				tree:      &matchersTree{},
				templates: []string{},
			},
			wantErr: true,
		},
		{
			name: "Test Case 4",
			args: args{
				tree:      &matchersTree{},
				templates: []string{"[a-z]+"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 5",
			args: args{
				tree:      &matchersTree{},
				templates: []string{"[a-z]+", "example.com"},
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

			if err := hostSNIRegexpV2(tt.args.tree, tt.args.templates...); (err != nil) != tt.wantErr {
				t.Errorf("hostSNIRegexpV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
