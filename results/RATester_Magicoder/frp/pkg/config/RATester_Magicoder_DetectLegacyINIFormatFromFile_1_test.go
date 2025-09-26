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
			name: "Test case 1",
			args: args{path: "testdata/valid.ini"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{path: "testdata/invalid.ini"},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{path: "testdata/non-existent.ini"},
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
