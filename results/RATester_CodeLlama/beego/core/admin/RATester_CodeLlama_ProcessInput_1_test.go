package admin

import (
	"fmt"
	"io"
	"os"
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
			name: "lookup goroutine",
			args: args{
				input: "lookup goroutine",
				w:     os.Stdout,
			},
		},
		{
			name: "lookup heap",
			args: args{
				input: "lookup heap",
				w:     os.Stdout,
			},
		},
		{
			name: "lookup threadcreate",
			args: args{
				input: "lookup threadcreate",
				w:     os.Stdout,
			},
		},
		{
			name: "lookup block",
			args: args{
				input: "lookup block",
				w:     os.Stdout,
			},
		},
		{
			name: "get cpuprof",
			args: args{
				input: "get cpuprof",
				w:     os.Stdout,
			},
		},
		{
			name: "get memprof",
			args: args{
				input: "get memprof",
				w:     os.Stdout,
			},
		},
		{
			name: "gc summary",
			args: args{
				input: "gc summary",
				w:     os.Stdout,
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
		})
	}
}
