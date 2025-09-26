package web

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestExecuteViewPathTemplate_1(t *testing.T) {
	type args struct {
		wr       io.Writer
		name     string
		viewPath string
		data     interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				wr:       &bytes.Buffer{},
				name:     "test.tmpl",
				viewPath: "views",
				data:     map[string]interface{}{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				wr:       &bytes.Buffer{},
				name:     "nonexistent.tmpl",
				viewPath: "views",
				data:     map[string]interface{}{"key": "value"},
			},
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				wr:       &bytes.Buffer{},
				name:     "test.tmpl",
				viewPath: "nonexistent_views",
				data:     map[string]interface{}{"key": "value"},
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

			if err := ExecuteViewPathTemplate(tt.args.wr, tt.args.name, tt.args.viewPath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ExecuteViewPathTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
