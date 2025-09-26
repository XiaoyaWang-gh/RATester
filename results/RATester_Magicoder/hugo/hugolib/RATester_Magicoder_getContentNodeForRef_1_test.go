package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestgetContentNodeForRef_1(t *testing.T) {
	type args struct {
		context      page.Page
		isReflink    bool
		hadExtension bool
		inRef        string
		ref          string
	}
	tests := []struct {
		name    string
		args    args
		want    contentNodeI
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
			got, err := c.getContentNodeForRef(tt.args.context, tt.args.isReflink, tt.args.hadExtension, tt.args.inRef, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("getContentNodeForRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContentNodeForRef() = %v, want %v", got, tt.want)
			}
		})
	}
}
