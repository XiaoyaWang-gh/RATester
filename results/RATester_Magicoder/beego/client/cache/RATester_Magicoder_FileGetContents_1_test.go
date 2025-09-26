package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileGetContents_1(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		want     []byte
		wantErr  bool
	}{
		{
			name:     "File Exists",
			filename: "testdata/test.txt",
			want:     []byte("test content"),
			wantErr:  false,
		},
		{
			name:     "File Does Not Exist",
			filename: "testdata/nonexistent.txt",
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := FileGetContents(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileGetContents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileGetContents() = %v, want %v", got, tt.want)
			}
		})
	}
}
