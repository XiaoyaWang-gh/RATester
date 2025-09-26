package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFilesWriter_1(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		{
			name: "Test newFilesWriter",
			want: &multiFileLogWriter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newFilesWriter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFilesWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
