package path

import (
	"fmt"
	"testing"
)

func TestDir_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		path    any
		want    string
		wantErr bool
	}{
		{
			name: "valid path",
			path: "/path/to/file.txt",
			want: "/path/to",
		},
		{
			name:    "invalid path",
			path:    123,
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

			got, err := ns.Dir(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}
