package tcp

import (
	"fmt"
	"testing"
)

func TestHostSNIRegexpV2_1(t *testing.T) {
	type args struct {
		templates []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid templates",
			args: args{
				templates: []string{"example.com", "^(www\\.)?.*$"},
			},
			wantErr: false,
		},
		{
			name: "invalid templates",
			args: args{
				templates: []string{"[example.com"},
			},
			wantErr: true,
		},
		{
			name: "empty templates",
			args: args{
				templates: []string{},
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

			tree := &matchersTree{}
			if err := hostSNIRegexpV2(tree, tt.args.templates...); (err != nil) != tt.wantErr {
				t.Errorf("hostSNIRegexpV2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
