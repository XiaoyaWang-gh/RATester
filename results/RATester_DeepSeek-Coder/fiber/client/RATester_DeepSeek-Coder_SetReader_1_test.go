package client

import (
	"fmt"
	"io"
	"testing"
)

func TestSetReader_1(t *testing.T) {
	type fields struct {
		reader    io.ReadCloser
		name      string
		fieldName string
		path      string
	}
	type args struct {
		r io.ReadCloser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			f := &File{
				reader:    tt.fields.reader,
				name:      tt.fields.name,
				fieldName: tt.fields.fieldName,
				path:      tt.fields.path,
			}
			f.SetReader(tt.args.r)
		})
	}
}
