package tcp

import (
	"fmt"
	"testing"
)

func TestHostSNIRegexp_1(t *testing.T) {
	type args struct {
		templates []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid template",
			args: args{
				templates: []string{"^www\\.example\\.com$"},
			},
			wantErr: false,
		},
		{
			name: "invalid template",
			args: args{
				templates: []string{"[a-z"},
			},
			wantErr: true,
		},
		{
			name: "non-ascii template",
			args: args{
				templates: []string{"日本語"},
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
			err := hostSNIRegexp(tree, tt.args.templates...)
			if (err != nil) != tt.wantErr {
				t.Errorf("hostSNIRegexp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
