package install

import (
	"fmt"
	"testing"
)

func TestPathExists_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				path: "path",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				path: "path",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test case 3",
			args: args{
				path: "path",
			},
			want:    true,
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

			got, err := pathExists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("pathExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
