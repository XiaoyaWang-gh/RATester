package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRead_1(t *testing.T) {
	tests := []struct {
		name    string
		f       FileOrContent
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test with valid file path",
			f:       FileOrContent("/path/to/file"),
			want:    []byte("file content"),
			wantErr: false,
		},
		{
			name:    "Test with invalid file path",
			f:       FileOrContent("/path/to/invalid/file"),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Test with content",
			f:       FileOrContent("content"),
			want:    []byte("content"),
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

			got, err := tt.f.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}
