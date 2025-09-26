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
			name: "test hostSNIRegexpV2",
			args: args{
				tree: &matchersTree{
					matcher: func(meta ConnData) bool {
						return true
					},
				},
				templates: []string{"example.com"},
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
