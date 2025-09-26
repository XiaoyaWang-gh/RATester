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
			name: "path",
			f:    FileOrContent("testdata/test.txt"),
			want: []byte("Hello, World!"),
		},
		{
			name: "content",
			f:    FileOrContent("Hello, World!"),
			want: []byte("Hello, World!"),
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
