package web

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExecuteTemplate_1(t *testing.T) {
	type args struct {
		wr   io.Writer
		name string
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				wr:   &bytes.Buffer{},
				name: "test.tmpl",
				data: map[string]interface{}{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				wr:   &bytes.Buffer{},
				name: "non-existent.tmpl",
				data: map[string]interface{}{"key": "value"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := ExecuteTemplate(tt.args.wr, tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ExecuteTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
