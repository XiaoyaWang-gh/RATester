package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"text/template"
)

func TestGetTemplate_1(t *testing.T) {
	type args struct {
		root   string
		fs     http.FileSystem
		file   string
		others []string
	}
	tests := []struct {
		name    string
		args    args
		wantT   *template.Template
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

			gotT, err := getTemplate(tt.args.root, tt.args.fs, tt.args.file, tt.args.others...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("getTemplate() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
