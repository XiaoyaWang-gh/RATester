package modules

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestRewriteGoMod_1(t *testing.T) {
	t.Parallel()

	type args struct {
		name    string
		isGoMod map[string]bool
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name:    "test.go",
				isGoMod: map[string]bool{"test.go": true},
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				name:    "test.go",
				isGoMod: map[string]bool{"test.go": false},
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

			c := &Client{
				fs: afero.NewMemMapFs(),
			}

			if err := c.fs.MkdirAll("test", 0755); err != nil {
				t.Fatal(err)
			}

			if err := afero.WriteFile(c.fs, "test/test.go", []byte("package main\nfunc main() {}\n"), 0644); err != nil {
				t.Fatal(err)
			}

			if err := c.rewriteGoMod(tt.args.name, tt.args.isGoMod); (err != nil) != tt.wantErr {
				t.Errorf("Client.rewriteGoMod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
