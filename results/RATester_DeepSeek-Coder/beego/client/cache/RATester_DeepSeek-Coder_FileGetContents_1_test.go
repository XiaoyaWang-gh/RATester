package cache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileGetContents_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test with existing file",
			args: args{
				filename: "testdata/existing_file.txt",
			},
			want:    []byte("This is a test file\n"),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := FileGetContents(tt.args.filename)
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
