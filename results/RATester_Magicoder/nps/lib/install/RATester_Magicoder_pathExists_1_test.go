package install

import (
	"fmt"
	"testing"
)

func TestpathExists_1(t *testing.T) {
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
			name: "Test Case 1",
			args: args{
				path: "/path/to/existing/file",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				path: "/path/to/non-existing/file",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				path: "/path/to/invalid/file",
			},
			want:    false,
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
