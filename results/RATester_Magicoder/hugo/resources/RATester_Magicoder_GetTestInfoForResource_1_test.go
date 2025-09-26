package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestGetTestInfoForResource_1(t *testing.T) {
	type args struct {
		r resource.Resource
	}
	tests := []struct {
		name string
		args args
		want GenericResourceTestInfo
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

			if got := GetTestInfoForResource(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestInfoForResource() = %v, want %v", got, tt.want)
			}
		})
	}
}
