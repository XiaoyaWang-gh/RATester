package npm

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestPack_1(t *testing.T) {
	type args struct {
		sourceFs                        afero.Fs
		assetsWithDuplicatesPreservedFs afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				sourceFs:                        &afero.MemMapFs{},
				assetsWithDuplicatesPreservedFs: &afero.MemMapFs{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				sourceFs:                        &afero.MemMapFs{},
				assetsWithDuplicatesPreservedFs: &afero.MemMapFs{},
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

			if err := Pack(tt.args.sourceFs, tt.args.assetsWithDuplicatesPreservedFs); (err != nil) != tt.wantErr {
				t.Errorf("Pack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
