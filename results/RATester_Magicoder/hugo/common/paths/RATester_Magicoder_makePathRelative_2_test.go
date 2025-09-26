package paths

import (
	"fmt"
	"testing"
)

func TestmakePathRelative_2(t *testing.T) {
	type args struct {
		inPath              string
		possibleDirectories []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				inPath:              "/path/to/file",
				possibleDirectories: []string{"/path", "/path/to"},
			},
			want:    "file",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				inPath:              "/path/to/file",
				possibleDirectories: []string{"/path", "/path/to/file"},
			},
			want:    "file",
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				inPath:              "/path/to/file",
				possibleDirectories: []string{"/path", "/path/to/dir"},
			},
			want:    "/path/to/file",
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

			got, err := makePathRelative(tt.args.inPath, tt.args.possibleDirectories...)
			if (err != nil) != tt.wantErr {
				t.Errorf("makePathRelative() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("makePathRelative() = %v, want %v", got, tt.want)
			}
		})
	}
}
