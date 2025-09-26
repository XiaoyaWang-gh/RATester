package hugofs

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestType_2(t *testing.T) {
	type fields struct {
		FileInfo fs.FileInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   fs.FileMode
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

			d := dirEntry{
				FileInfo: tt.fields.FileInfo,
			}
			if got := d.Type(); got != tt.want {
				t.Errorf("dirEntry.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}
