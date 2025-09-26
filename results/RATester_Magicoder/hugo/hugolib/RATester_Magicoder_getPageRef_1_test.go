package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestgetPageRef_1(t *testing.T) {
	type args struct {
		context page.Page
		ref     string
	}
	tests := []struct {
		name    string
		args    args
		want    page.Page
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

			c := &pageFinder{}
			got, err := c.getPageRef(tt.args.context, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPageRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPageRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
