package admin

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestMemProf_1(t *testing.T) {
	tests := []struct {
		name string
		w    io.Writer
	}{
		{
			name: "Test case 1",
			w:    &bytes.Buffer{},
		},
		{
			name: "Test case 2",
			w:    os.Stdout,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			MemProf(tt.w)
			// Check if the file was created
			_, err := os.Stat("mem-" + strconv.Itoa(pid) + ".memprof")
			if err != nil {
				t.Errorf("Expected file to be created, but got error: %v", err)
			}
			// Check if the message was written to the writer
			expectedMessage := fmt.Sprintf("create heap profile mem-%s.memprof \nNow you can use this to check it: go tool pprof %s mem-%s.memprof\n", strconv.Itoa(pid), os.Args[0], strconv.Itoa(pid))
			if tt.w.(*bytes.Buffer).String() != expectedMessage {
				t.Errorf("Expected message to be %s, but got %s", expectedMessage, tt.w.(*bytes.Buffer).String())
			}
		})
	}
}
