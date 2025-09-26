package echo

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestStaticFileHandler_1(t *testing.T) {
	type args struct {
		file       string
		filesystem fs.FS
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := StaticFileHandler(tt.args.file, tt.args.filesystem)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StaticFileHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
