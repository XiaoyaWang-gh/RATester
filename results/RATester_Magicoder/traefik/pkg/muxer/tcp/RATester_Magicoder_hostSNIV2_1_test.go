package tcp

import (
	"fmt"
	"testing"
)

func TestHostSNIV2_1(t *testing.T) {
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
				hosts: []string{"example.com", "*.example.com"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{},
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"*"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"invalid_hostname"},
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

			if err := hostSNIV2(tt.args.tree, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("hostSNIV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
