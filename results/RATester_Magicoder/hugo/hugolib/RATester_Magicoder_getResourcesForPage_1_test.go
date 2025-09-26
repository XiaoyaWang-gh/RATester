package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestgetResourcesForPage_1(t *testing.T) {
	type args struct {
		ps *pageState
	}
	tests := []struct {
		name    string
		args    args
		want    resource.Resources
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

			m := &pageMap{}
			got, err := m.getResourcesForPage(tt.args.ps)
			if (err != nil) != tt.wantErr {
				t.Errorf("getResourcesForPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResourcesForPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
