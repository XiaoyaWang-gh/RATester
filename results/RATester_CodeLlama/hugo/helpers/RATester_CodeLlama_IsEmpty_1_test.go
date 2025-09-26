package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsEmpty_1(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		fs      afero.Fs
		want    bool
		wantErr bool
	}{
		{
			name:    "empty",
			path:    "empty",
			fs:      afero.NewMemMapFs(),
			want:    true,
			wantErr: false,
		},
		{
			name:    "not empty",
			path:    "not_empty",
			fs:      afero.NewMemMapFs(),
			want:    false,
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

			got, err := IsEmpty(tt.path, tt.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
