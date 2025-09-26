package config

import (
	"fmt"
	"testing"
)

func TestDetectLegacyINIFormatFromFile_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test with existing file",
			args: args{path: "testdata/existing_file.ini"},
			want: true,
		},
		{
			name: "Test with non-existing file",
			args: args{path: "testdata/non_existing_file.ini"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DetectLegacyINIFormatFromFile(tt.args.path); got != tt.want {
				t.Errorf("DetectLegacyINIFormatFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
