package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLogProject_1(t *testing.T) {
	type args struct {
		name            string
		endpoint        string
		AccessKeyID     string
		accessKeySecret string
	}
	tests := []struct {
		name    string
		args    args
		want    *LogProject
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				name:            "test_project",
				endpoint:        "http://localhost:8080",
				AccessKeyID:     "test_key_id",
				accessKeySecret: "test_key_secret",
			},
			want: &LogProject{
				Name:            "test_project",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test_key_id",
				AccessKeySecret: "test_key_secret",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewLogProject(tt.args.name, tt.args.endpoint, tt.args.AccessKeyID, tt.args.accessKeySecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLogProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogProject() = %v, want %v", got, tt.want)
			}
		})
	}
}
