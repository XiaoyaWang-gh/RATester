package admin

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestProcessInput_1(t *testing.T) {
	type args struct {
		input string
		w     io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test lookup goroutine",
			args: args{
				input: "lookup goroutine",
				w:     &bytes.Buffer{},
			},
		},
		{
			name: "Test lookup heap",
			args: args{
				input: "lookup heap",
				w:     &bytes.Buffer{},
			},
		},
		{
			name: "Test lookup threadcreate",
			args: args{
				input: "lookup threadcreate",
				w:     &bytes.Buffer{},
			},
		},
		{
			name: "Test lookup block",
			args: args{
				input: "lookup block",
				w:     &bytes.Buffer{},
			},
		},
		{
			name: "Test get cpuprof",
			args: args{
				input: "get cpuprof",
				w:     &bytes.Buffer{},
			},
		},
		{
			name: "Test get memprof",
			args: args{
				input: "get memprof",
				w:     &bytes.Buffer{},
			},
		},
		{
			name: "Test gc summary",
			args: args{
				input: "gc summary",
				w:     &bytes.Buffer{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ProcessInput(tt.args.input, tt.args.w)
			// Add assertions here to check the output of the function
		})
	}
}
