package skeletons

import (
	"embed"
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCopyFiles_1(t *testing.T) {
	type args struct {
		createpath string
		sourceFs   afero.Fs
		skeleton   embed.FS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test copyFiles",
			args: args{
				createpath: "test",
				sourceFs:   afero.NewMemMapFs(),
				skeleton:   embed.FS{},
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

			if err := copyFiles(tt.args.createpath, tt.args.sourceFs, tt.args.skeleton); (err != nil) != tt.wantErr {
				t.Errorf("copyFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
