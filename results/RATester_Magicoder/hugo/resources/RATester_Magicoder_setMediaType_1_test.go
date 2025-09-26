package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestsetMediaType_1(t *testing.T) {
	type args struct {
		mediaType media.Type
	}
	tests := []struct {
		name string
		l    *genericResource
		args args
		want media.Type
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

			tt.l.setMediaType(tt.args.mediaType)
			if got := tt.l.sd.MediaType; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.setMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}
