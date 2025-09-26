package commands

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestReaddir_3(t *testing.T) {
	type testCase struct {
		name    string
		f       noDirFile
		count   int
		want    []fs.FileInfo
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "Test case 1",
			f:       noDirFile{},
			count:   10,
			want:    nil,
			wantErr: false,
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

			got, err := tt.f.Readdir(tt.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("noDirFile.Readdir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noDirFile.Readdir() = %v, want %v", got, tt.want)
			}
		})
	}
}
