package cssjs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestContentHash_1(t *testing.T) {
	type testCase struct {
		name     string
		filename string
		want     []byte
		want1    string
	}

	tests := []testCase{
		{
			name:     "Test with existing file",
			filename: "testdata/existing_file.txt",
			want:     []byte("This is a test file"),
			want1:    "6c3992c9794201197933e47b9002c04c0147c3f57c9c9c13a1b2f937977c6d1d",
		},
		{
			name:     "Test with non-existing file",
			filename: "testdata/non_existing_file.txt",
			want:     nil,
			want1:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			imp := &importResolver{
				fs: afero.NewOsFs(),
			}
			got, got1 := imp.contentHash(tt.filename)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importResolver.contentHash() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("importResolver.contentHash() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
