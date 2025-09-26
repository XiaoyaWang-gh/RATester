package path

import (
	"fmt"
	"testing"
)

func TestBase_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		path    any
		want    string
		wantErr bool
	}{
		{
			name:    "valid path",
			path:    "/foo/bar/baz.txt",
			want:    "baz.txt",
			wantErr: false,
		},
		{
			name:    "invalid path",
			path:    123,
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty path",
			path:    "",
			want:    "",
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

			got, err := ns.Base(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base() = %v, want %v", got, tt.want)
			}
		})
	}
}
