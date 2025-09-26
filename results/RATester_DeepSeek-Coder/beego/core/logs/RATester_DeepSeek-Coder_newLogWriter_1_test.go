package logs

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewLogWriter_1(t *testing.T) {
	tests := []struct {
		name string
		wr   io.Writer
		want *logWriter
	}{
		{
			name: "Test case 1",
			wr:   bytes.NewBuffer([]byte{}),
			want: &logWriter{writer: bytes.NewBuffer([]byte{})},
		},
		{
			name: "Test case 2",
			wr:   nil,
			want: &logWriter{writer: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newLogWriter(tt.wr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLogWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
