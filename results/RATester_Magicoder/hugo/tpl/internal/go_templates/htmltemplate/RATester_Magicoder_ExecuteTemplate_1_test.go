package template

import (
	"fmt"
	"io"
	"testing"
)

func TestExecuteTemplate_1(t *testing.T) {
	type args struct {
		wr   io.Writer
		name string
		data any
	}
	tests := []struct {
		name    string
		t       *Template
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

			if err := tt.t.ExecuteTemplate(tt.args.wr, tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Template.ExecuteTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
