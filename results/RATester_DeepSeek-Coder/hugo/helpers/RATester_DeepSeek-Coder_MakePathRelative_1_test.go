package helpers

import (
	"fmt"
	"testing"
)

func TestMakePathRelative_1(t *testing.T) {
	type test struct {
		inPath              string
		possibleDirectories []string
		want                string
		wantErr             bool
	}

	tests := []test{
		{
			inPath:              "/home/user/file.txt",
			possibleDirectories: []string{"/home/user", "/var/lib"},
			want:                "file.txt",
			wantErr:             false,
		},
		{
			inPath:              "/var/lib/file.txt",
			possibleDirectories: []string{"/home/user", "/var/lib"},
			want:                "file.txt",
			wantErr:             false,
		},
		{
			inPath:              "/var/lib/file.txt",
			possibleDirectories: []string{"/home/user", "/tmp"},
			want:                "/var/lib/file.txt",
			wantErr:             true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("inPath=%s,possibleDirectories=%s", tt.inPath, tt.possibleDirectories), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := MakePathRelative(tt.inPath, tt.possibleDirectories...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakePathRelative() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakePathRelative() = %v, want %v", got, tt.want)
			}
		})
	}
}
