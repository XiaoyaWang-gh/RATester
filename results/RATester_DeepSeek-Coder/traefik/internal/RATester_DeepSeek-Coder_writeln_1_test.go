package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestWriteln_1(t *testing.T) {
	tests := []struct {
		name string
		w    io.Writer
		a    []interface{}
		want error
	}{
		{
			name: "Test with nil writer",
			w:    nil,
			a:    []interface{}{"test"},
			want: fmt.Errorf("io: write on closed writer"),
		},
		{
			name: "Test with valid writer",
			w:    os.Stdout,
			a:    []interface{}{"test"},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ew := &errWriter{
				w: tt.w,
			}

			ew.writeln(tt.a...)

			if !errors.Is(ew.err, tt.want) {
				t.Errorf("got %v, want %v", ew.err, tt.want)
			}
		})
	}
}
