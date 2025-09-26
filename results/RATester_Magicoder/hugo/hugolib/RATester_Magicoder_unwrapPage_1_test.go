package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestunwrapPage_1(t *testing.T) {
	type args struct {
		in any
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

			got, err := unwrapPage(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("unwrapPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unwrapPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
