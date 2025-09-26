package create

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestMatch_1(t *testing.T) {
	type args struct {
		name      string
		pattern   string
		matchFunc func(r resource.Resource) bool
		firstOnly bool
	}
	tests := []struct {
		name    string
		c       *Client
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

			got, err := tt.c.match(tt.args.name, tt.args.pattern, tt.args.matchFunc, tt.args.firstOnly)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
