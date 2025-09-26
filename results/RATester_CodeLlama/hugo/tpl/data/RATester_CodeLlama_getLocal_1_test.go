package data

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestGetLocal_1(t *testing.T) {
	type args struct {
		workingDir string
		url        string
		fs         afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				workingDir: "",
				url:        "",
				fs:         nil,
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

			got, err := getLocal(tt.args.workingDir, tt.args.url, tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
