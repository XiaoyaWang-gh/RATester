package web

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"testing"
)

func TestGetTplDeep_1(t *testing.T) {
	type args struct {
		root   string
		fs     http.FileSystem
		file   string
		parent string
		t      *template.Template
	}
	tests := []struct {
		name    string
		args    args
		want    *template.Template
		want1   [][]string
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

			got, got1, err := getTplDeep(tt.args.root, tt.args.fs, tt.args.file, tt.args.parent, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTplDeep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTplDeep() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getTplDeep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
