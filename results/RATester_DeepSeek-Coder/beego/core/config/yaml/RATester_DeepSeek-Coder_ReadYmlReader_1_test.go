package yaml

import (
	"fmt"
	"testing"
)

func TestReadYmlReader_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with existing file",
			args: args{
				path: "testdata/existing_file.yml",
			},
			wantErr: false,
		},
		{
			name: "Test with non-existing file",
			args: args{
				path: "testdata/non_existing_file.yml",
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

			_, err := ReadYmlReader(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadYmlReader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
