package accesslog

import (
	"fmt"
	"testing"
)

func TestOpenAccessLogFile_1(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid file path",
			args: args{
				filePath: "/tmp/test.log",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid file path",
			args: args{
				filePath: "/invalid/path",
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

			_, err := openAccessLogFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("openAccessLogFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
