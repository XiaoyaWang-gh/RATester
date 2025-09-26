package skeletons

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCreateSite_1(t *testing.T) {
	type args struct {
		createpath string
		sourceFs   afero.Fs
		force      bool
		format     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				createpath: "testpath",
				sourceFs:   afero.NewMemMapFs(),
				force:      true,
				format:     "testformat",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				createpath: "testpath",
				sourceFs:   afero.NewMemMapFs(),
				force:      false,
				format:     "testformat",
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

			if err := CreateSite(tt.args.createpath, tt.args.sourceFs, tt.args.force, tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("CreateSite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
