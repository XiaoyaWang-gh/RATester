package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseFiles_1(t *testing.T) {
	type args struct {
		t         *Template
		readFile  func(string) (string, []byte, error)
		filenames []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Template
		wantErr bool
	}{
		{
			name: "test_case_1",
			args: args{
				t: nil,
				readFile: func(string) (string, []byte, error) {
					return "", nil, nil
				},
				filenames: nil,
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

			got, err := parseFiles(tt.args.t, tt.args.readFile, tt.args.filenames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}
