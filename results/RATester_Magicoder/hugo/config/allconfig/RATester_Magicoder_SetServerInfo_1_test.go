package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/urls"
)

func TestSetServerInfo_1(t *testing.T) {
	type args struct {
		baseURL           urls.BaseURL
		baseURLLiveReload urls.BaseURL
		serverInterface   string
	}
	tests := []struct {
		name string
		args args
		want ConfigCompiled
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

			c := &ConfigCompiled{}
			c.SetServerInfo(tt.args.baseURL, tt.args.baseURLLiveReload, tt.args.serverInterface)
			if !reflect.DeepEqual(*c, tt.want) {
				t.Errorf("SetServerInfo() = %v, want %v", *c, tt.want)
			}
		})
	}
}
