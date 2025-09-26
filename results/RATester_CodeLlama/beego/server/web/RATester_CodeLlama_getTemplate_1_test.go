package web

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"testing"
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
		want    *template.Template
		wantErr bool
	}{
		{
			"test_getTemplate_001",
			args{
				root:   "root",
				fs:     http.Dir("fs"),
				file:   "file",
				others: []string{"others"},
			},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := getTemplate(tt.args.root, tt.args.fs, tt.args.file, tt.args.others...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTemplate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
