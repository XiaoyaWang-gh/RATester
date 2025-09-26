package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/doctree"
)

func TestforEachResourceInPage_1(t *testing.T) {
	type args struct {
		ps       *pageState
		lockType doctree.LockType
		exact    bool
		handle   func(resourceKey string, n contentNodeI, match doctree.DimensionFlag) (bool, error)
	}
	tests := []struct {
		name    string
		m       *pageMap
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

			if err := tt.m.forEachResourceInPage(tt.args.ps, tt.args.lockType, tt.args.exact, tt.args.handle); (err != nil) != tt.wantErr {
				t.Errorf("forEachResourceInPage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
