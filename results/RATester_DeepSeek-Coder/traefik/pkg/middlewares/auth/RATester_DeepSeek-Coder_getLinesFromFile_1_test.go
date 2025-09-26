package auth

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLinesFromFile_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test with existing file",
			args: args{
				filename: "testdata/existing_file.txt",
			},
			want:    []string{"line1", "line2", "line3"},
			wantErr: false,
		},
		{
			name: "Test with non-existing file",
			args: args{
				filename: "testdata/non_existing_file.txt",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test with empty file",
			args: args{
				filename: "testdata/empty_file.txt",
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "Test with file containing only comments",
			args: args{
				filename: "testdata/comments_only_file.txt",
			},
			want:    []string{},
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

			got, err := getLinesFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLinesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLinesFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
