package filesystems

import (
	"fmt"
	"testing"
)

func TestRealFilename_1(t *testing.T) {
	type args struct {
		rel string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &SourceFilesystem{}
			if got := d.RealFilename(tt.args.rel); got != tt.want {
				t.Errorf("RealFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
