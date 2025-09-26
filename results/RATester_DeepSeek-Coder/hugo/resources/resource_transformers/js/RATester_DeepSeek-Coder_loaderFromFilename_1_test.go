package js

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/evanw/esbuild/pkg/api"
)

func TestLoaderFromFilename_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     api.Loader
	}{
		{
			name:     "Test .js file",
			filename: "test.js",
			want:     api.LoaderJS,
		},
		{
			name:     "Test .css file",
			filename: "test.css",
			want:     api.LoaderCSS,
		},
		{
			name:     "Test .json file",
			filename: "test.json",
			want:     api.LoaderJSON,
		},
		{
			name:     "Test .txt file",
			filename: "test.txt",
			want:     api.LoaderJS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := loaderFromFilename(tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loaderFromFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
