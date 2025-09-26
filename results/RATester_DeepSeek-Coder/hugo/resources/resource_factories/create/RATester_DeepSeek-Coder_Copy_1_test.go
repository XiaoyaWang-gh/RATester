package create

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestCopy_1(t *testing.T) {
	type args struct {
		r          resource.Resource
		targetPath string
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

			got, err := tt.c.Copy(tt.args.r, tt.args.targetPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
