package skeletons

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCreateTheme_1(t *testing.T) {
	type args struct {
		createpath string
		sourceFs   afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				createpath: "test",
				sourceFs:   afero.NewMemMapFs(),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				createpath: "test",
				sourceFs:   afero.NewOsFs(),
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

			err := CreateTheme(tt.args.createpath, tt.args.sourceFs)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTheme() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
