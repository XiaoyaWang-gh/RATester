package web

import (
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
		// TODO: Add test cases.
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
