package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestNewIdentityManager_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    ConfigLanguage
		args args
		want identity.Manager
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

			if got := tt.c.NewIdentityManager(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentityManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
