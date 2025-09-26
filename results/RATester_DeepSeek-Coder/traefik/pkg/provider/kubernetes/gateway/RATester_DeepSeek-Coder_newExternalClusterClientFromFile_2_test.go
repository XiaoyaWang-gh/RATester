package gateway

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewExternalClusterClientFromFile_2(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *clientWrapper
		wantErr bool
	}{
		{
			name: "Test with valid kubeconfig file",
			args: args{
				file: "testdata/valid_kubeconfig.yaml",
			},
			want:    &clientWrapper{},
			wantErr: false,
		},
		{
			name: "Test with invalid kubeconfig file",
			args: args{
				file: "testdata/invalid_kubeconfig.yaml",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test with non-existent kubeconfig file",
			args: args{
				file: "testdata/non_existent_kubeconfig.yaml",
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

			got, err := newExternalClusterClientFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("newExternalClusterClientFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newExternalClusterClientFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
