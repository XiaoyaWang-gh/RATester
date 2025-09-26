package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/source"
)

func TesthandleDataFile_1(t *testing.T) {
	type args struct {
		r *source.File
	}
	tests := []struct {
		name    string
		h       *HugoSites
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

			if err := tt.h.handleDataFile(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("HugoSites.handleDataFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
