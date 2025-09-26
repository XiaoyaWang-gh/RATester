package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMode_1(t *testing.T) {
	type fields struct {
		name    string
		size    int64
		mode    os.FileMode
		modTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   os.FileMode
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

			fi := bindataFileInfo{
				name:    tt.fields.name,
				size:    tt.fields.size,
				mode:    tt.fields.mode,
				modTime: tt.fields.modTime,
			}
			if got := fi.Mode(); got != tt.want {
				t.Errorf("bindataFileInfo.Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}
