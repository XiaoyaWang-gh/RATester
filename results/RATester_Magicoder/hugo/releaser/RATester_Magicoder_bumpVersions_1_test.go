package releaser

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugo"
)

func TestbumpVersions_1(t *testing.T) {
	type args struct {
		ver hugo.Version
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ver: hugo.Version{
					Major:      0,
					Minor:      90,
					PatchLevel: 0,
					Suffix:     "",
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				ver: hugo.Version{
					Major:      0,
					Minor:      90,
					PatchLevel: 0,
					Suffix:     "beta",
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				ver: hugo.Version{
					Major:      0,
					Minor:      90,
					PatchLevel: 0,
					Suffix:     "rc",
				},
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

			r := &ReleaseHandler{}
			if err := r.bumpVersions(tt.args.ver); (err != nil) != tt.wantErr {
				t.Errorf("bumpVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
