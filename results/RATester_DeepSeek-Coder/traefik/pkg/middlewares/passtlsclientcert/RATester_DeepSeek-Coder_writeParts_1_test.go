package passtlsclientcert

import (
	"context"
	"fmt"
	"io"
	"testing"
)

func TestWriteParts_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		content io.StringWriter
		entries []string
		prefix  string
	}
	tests := []struct {
		name string
		args args
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

			writeParts(tt.args.ctx, tt.args.content, tt.args.entries, tt.args.prefix)
		})
	}
}
