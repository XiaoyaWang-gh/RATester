package echo

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestStaticFileHandler_1(t *testing.T) {
	type args struct {
		file       string
		filesystem fs.FS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				file:       "test.txt",
				filesystem: fstest.MapFS{"test.txt": &fstest.MapFile{Data: []byte("test data")}},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				file:       "non-existent.txt",
				filesystem: fstest.MapFS{},
			},
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

			handler := StaticFileHandler(tt.args.file, tt.args.filesystem)
			if err := handler(nil); (err != nil) != tt.wantErr {
				t.Errorf("StaticFileHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
