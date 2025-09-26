package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilepathOrContent_1(t *testing.T) {
	tests := []struct {
		name          string
		fileOrContent interface{}
		wantContent   []byte
		wantErr       bool
	}{
		{
			name:          "Test with valid file path",
			fileOrContent: "testdata/test.txt",
			wantContent:   []byte("test content"),
			wantErr:       false,
		},
		{
			name:          "Test with valid byte content",
			fileOrContent: []byte("test content"),
			wantContent:   []byte("test content"),
			wantErr:       false,
		},
		{
			name:          "Test with invalid file path",
			fileOrContent: "testdata/invalid.txt",
			wantContent:   nil,
			wantErr:       true,
		},
		{
			name:          "Test with invalid type",
			fileOrContent: 123,
			wantContent:   nil,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotContent, err := filepathOrContent(tt.fileOrContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("filepathOrContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotContent, tt.wantContent) {
				t.Errorf("filepathOrContent() = %v, want %v", gotContent, tt.wantContent)
			}
		})
	}
}
