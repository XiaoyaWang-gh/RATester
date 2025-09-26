package integrity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestFingerprint_2(t *testing.T) {
	type args struct {
		res  resources.ResourceTransformer
		algo string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    resource.Resource
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

			got, err := tt.c.Fingerprint(tt.args.res, tt.args.algo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Fingerprint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Fingerprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
