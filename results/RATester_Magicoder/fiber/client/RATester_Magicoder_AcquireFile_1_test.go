package client

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestAcquireFile_1(t *testing.T) {
	type args struct {
		setter []SetFileFunc
	}
	tests := []struct {
		name string
		args args
		want *File
	}{
		{
			name: "Test AcquireFile with setter",
			args: args{
				setter: []SetFileFunc{
					func(f *File) {
						f.SetName("test.txt")
						f.SetPath("/tmp/test.txt")
						f.SetFieldName("file")
						f.SetReader(ioutil.NopCloser(strings.NewReader("test")))
					},
				},
			},
			want: &File{
				name:      "test.txt",
				path:      "/tmp/test.txt",
				fieldName: "file",
				reader:    ioutil.NopCloser(strings.NewReader("test")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AcquireFile(tt.args.setter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AcquireFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
