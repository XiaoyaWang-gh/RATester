package os

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestReadDir_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		input     any
		want      []fs.FileInfo
		wantError bool
	}{
		{
			name:      "success",
			input:     "/path/to/dir",
			want:      []fs.FileInfo{},
			wantError: false,
		},
		{
			name:      "failure",
			input:     "/path/to/non-existent/dir",
			want:      nil,
			wantError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			ns := &Namespace{
				workFs: afero.NewMemMapFs(),
			}

			got, err := ns.ReadDir(tc.input)

			if (err != nil) != tc.wantError {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tc.wantError)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReadDir() got = %v, want %v", got, tc.want)
			}
		})
	}
}
