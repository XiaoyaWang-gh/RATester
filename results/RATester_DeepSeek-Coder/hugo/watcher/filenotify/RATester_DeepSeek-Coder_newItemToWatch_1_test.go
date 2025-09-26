package filenotify

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewItemToWatch_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *itemToWatch
		wantErr bool
	}{
		{
			name: "Test with existing file",
			args: args{filename: "testdata/existing_file.txt"},
			want: &itemToWatch{
				filename: "testdata/existing_file.txt",
				left:     &recording{},
			},
			wantErr: false,
		},
		{
			name:    "Test with non-existing file",
			args:    args{filename: "testdata/non_existing_file.txt"},
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

			got, err := newItemToWatch(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("newItemToWatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newItemToWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
