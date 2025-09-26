package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestgetContentNodeFromRefReverseLookup_1(t *testing.T) {
	type args struct {
		ref string
		fi  hugofs.FileMetaInfo
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
			got, err := c.getContentNodeFromRefReverseLookup(tt.args.ref, tt.args.fi)
			if (err != nil) != tt.wantErr {
				t.Errorf("getContentNodeFromRefReverseLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContentNodeFromRefReverseLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
