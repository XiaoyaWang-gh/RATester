package i18n

import (
	"fmt"
	"testing"

	"github.com/gohugoio/go-i18n/v2/i18n"
	"github.com/gohugoio/hugo/source"
)

func TestaddTranslationFile_1(t *testing.T) {
	type args struct {
		bundle *i18n.Bundle
		r      *source.File
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

			if err := addTranslationFile(tt.args.bundle, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("addTranslationFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
