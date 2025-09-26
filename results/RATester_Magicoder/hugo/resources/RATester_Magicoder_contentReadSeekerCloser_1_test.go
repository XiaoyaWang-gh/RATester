package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestcontentReadSeekerCloser_1(t *testing.T) {
	type args struct {
		r resource.Resource
	}
	tests := []struct {
		name    string
		args    args
		want    hugio.ReadSeekCloser
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

			got, err := contentReadSeekerCloser(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("contentReadSeekerCloser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("contentReadSeekerCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
