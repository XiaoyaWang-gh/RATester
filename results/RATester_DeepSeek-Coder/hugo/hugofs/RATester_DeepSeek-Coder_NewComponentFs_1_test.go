package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestNewComponentFs_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		opts    ComponentFsOptions
		wantErr bool
	}{
		{
			name: "valid options",
			opts: ComponentFsOptions{
				Fs:                     afero.NewMemMapFs(),
				Component:              "content",
				DefaultContentLanguage: "en",
			},
			wantErr: false,
		},
		{
			name: "missing Fs",
			opts: ComponentFsOptions{
				Component:              "content",
				DefaultContentLanguage: "en",
			},
			wantErr: true,
		},
		{
			name: "missing Component",
			opts: ComponentFsOptions{
				Fs:                     afero.NewMemMapFs(),
				DefaultContentLanguage: "en",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); (r != nil) != tc.wantErr {
					t.Errorf("NewComponentFs() = %v, wantErr %v", r, tc.wantErr)
				}
			}()
			NewComponentFs(tc.opts)
		})
	}
}
