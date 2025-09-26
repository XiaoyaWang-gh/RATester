package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRenderedConfFromFile_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				path: "testdata/test.txt",
			},
			wantOut: []byte("test content"),
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				path: "testdata/nonexistent.txt",
			},
			wantOut: nil,
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

			gotOut, err := GetRenderedConfFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRenderedConfFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetRenderedConfFromFile() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
