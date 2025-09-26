package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestAddTemplateFile_1(t *testing.T) {
	type args struct {
		name string
		fim  hugofs.FileMetaInfo
	}
	tests := []struct {
		name    string
		t       *templateHandler
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

			if err := tt.t.addTemplateFile(tt.args.name, tt.args.fim); (err != nil) != tt.wantErr {
				t.Errorf("templateHandler.addTemplateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
