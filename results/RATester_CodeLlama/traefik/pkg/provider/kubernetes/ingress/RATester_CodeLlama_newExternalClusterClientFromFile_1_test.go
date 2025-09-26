package ingress

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewExternalClusterClientFromFile_1(t *testing.T) {
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
			name: "test",
			args: args{
				file: "test",
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
				t.Errorf("newExternalClusterClientFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
